package httpadapter

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/getsentry/raven-go"
	"github.com/gorilla/mux"
	"github.com/republicprotocol/renex-ingress-go/ingress"
	"github.com/republicprotocol/swapperd/foundation"
	"github.com/rs/cors"
	"golang.org/x/time/rate"
)

type loginRequest struct {
	Address  string `json:"address"`
	Referrer string `json:"referrer"`
}

type loginResponse struct {
	Verified bool `json:"verified"`
}

type kyberRequest struct {
	Address string            `json:"address"`
	Request clientAuthRequest `json:"request"`
}

type appAuthRequest struct {
	Type         string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type clientAuthRequest struct {
	Type   string `json:"grant_type"`
	Code   string `json:"code"`
	URI    string `json:"redirect_uri"`
	Key    string `json:"client_id"`
	Secret string `json:"client_secret"`
}

type tokenResponse struct {
	Type         string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	Expiry       int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type userResponse struct {
	UID       int      `json:"uid"`
	Status    string   `json:"kyc_status"`
	Addresses []string `json:"active_wallets"`
}

type usersResponse struct {
	Users []userResponse `json:"authorized_users"`
}

type delayInfo struct {
	OrderID string `json:"order_id"`
	KycAddr string `json:"kyc_addr"`
}

const (
	statusApproved string = "approved"
	statusPending  string = "pending"
	statusNone     string = "none"
)

// NewIngressServer returns an http server that forwards requests to an
// IngressAdapter.
func NewIngressServer(ingressAdapter IngressAdapter, approvedTraders []string, kyberID, kyberSecret string) http.Handler {
	limiter := rate.NewLimiter(3, 20)
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/orders", rateLimit(limiter, OpenOrderHandler(ingressAdapter, approvedTraders, kyberID, kyberSecret))).Methods("POST")
	r.HandleFunc("/login", rateLimit(limiter, LoginHandler(ingressAdapter, kyberID, kyberSecret))).Methods("POST")
	r.HandleFunc("/kyber", rateLimit(limiter, KyberKYCHandler(ingressAdapter, kyberID, kyberSecret))).Methods("POST")
	r.HandleFunc("/withdrawals", rateLimit(limiter, ApproveWithdrawalHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/address", rateLimit(limiter, PostAddressHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/swap", rateLimit(limiter, PostSwapHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/authorize", rateLimit(limiter, PostAuthorizeHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/kyc/{address}", rateLimit(limiter, GetKYCHandler(ingressAdapter, kyberID, kyberSecret))).Methods("GET")
	r.HandleFunc("/authorized/{address}", rateLimit(limiter, GetAuthorizedHandler(ingressAdapter))).Methods("GET")
	r.HandleFunc("/address/{orderID}", rateLimit(limiter, GetAddressHandler(ingressAdapter))).Methods("GET")
	r.HandleFunc("/swap/{orderID}", rateLimit(limiter, GetSwapHandler(ingressAdapter))).Methods("GET")
	r.HandleFunc("/swapperd/cb", rateLimit(limiter, PostSwapCallbackHandler(ingressAdapter))).Methods("POST")
	r.Use(RecoveryHandler)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)

	return handler
}

// OpenOrderHandler handles all HTTP open order requests
func OpenOrderHandler(ingressAdapter IngressAdapter, approvedTraders []string, kyberID, kyberSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		openOrderRequest := OpenOrderRequest{}
		if err := json.NewDecoder(r.Body).Decode(&openOrderRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode json into an order or a list of order fragments: %v", err)))
			return
		}

		// If the trader has not been manually approved (e.g. Lotan traders),
		// check their verification status.
		if !traderApproved(openOrderRequest.Address, approvedTraders) {
			kycType, err := traderVerified(ingressAdapter, kyberID, kyberSecret, openOrderRequest.Address)
			if err != nil {
				errString := fmt.Sprintf("cannot check trader verification: %v", err)
				log.Println(errString)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errString))
				raven.CaptureErrorAndWait(errors.New(errString), map[string]string{
					"trader": openOrderRequest.Address,
				})
				return
			}
			if kycType == ingress.KYCNone {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("trader is not verified"))
				return
			}
		}

		signature, err := ingressAdapter.OpenOrder(openOrderRequest.Address, openOrderRequest.OrderFragmentMappings)
		if err != nil {
			errString := fmt.Sprintf("cannot open order: %v", err)
			log.Println(errString)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errString))
			raven.CaptureErrorAndWait(errors.New(errString), nil)
			return
		}

		response, err := json.Marshal(OpenOrderResponse{
			Signature: MarshalSignature(signature),
		})
		if err != nil {
			errString := fmt.Sprintf("cannot open order: %v", err)
			log.Println(errString)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errString))
			raven.CaptureErrorAndWait(errors.New(errString), nil)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

// LoginHandler handles trader login requests
func LoginHandler(loginAdapter LoginAdapter, kyberID, kyberSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode POST request data
		decoder := json.NewDecoder(r.Body)
		var data loginRequest
		err := decoder.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode data: %v", err)))
			return
		}

		// Store address in database if it does not already exist
		if err := loginAdapter.PostLogin(data.Address, data.Referrer); err != nil {
			errString := fmt.Sprintf("cannot store login address: %v", err)
			log.Println(errString)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errString))
			raven.CaptureErrorAndWait(errors.New(errString), map[string]string{
				"trader": data.Address,
			})
			return
		}

		// Check if the trader is verified
		kycType, err := traderVerified(loginAdapter, kyberID, kyberSecret, data.Address)
		if err != nil {
			errString := fmt.Sprintf("cannot check trader verification: %v", err)
			log.Println(errString)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errString))
			raven.CaptureErrorAndWait(errors.New(errString), map[string]string{
				"trader": data.Address,
			})
			return
		}

		response := loginResponse{
			Verified: kycType != ingress.KYCNone,
		}
		respBytes, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot marshal login response: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	}
}

// KyberKYCHandler handles all Kyber authorization requests
func KyberKYCHandler(loginAdapter LoginAdapter, kyberID, kyberSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode POST request data
		decoder := json.NewDecoder(r.Body)
		var data kyberRequest
		err := decoder.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode data: %v", err)))
			return
		}

		// Construct new request object with Kyber secret key
		authRequest := data.Request
		authRequest.Secret = kyberSecret
		byteArray, err := json.Marshal(authRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot marshal data: %v", err)))
			return
		}

		// Forward updated request data to Kyber
		url := "https://kyber.network/oauth/token"
		postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(byteArray))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot construct new request: %v", err)))
			return
		}
		postRequest.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(postRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("unable to forward request: %v", err)))
			return
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("unable to read kyber response: %v", err)))
			return
		}

		if len(bodyBytes) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid authorization code"))
			return
		}

		var tokenResp tokenResponse
		if err := json.Unmarshal(bodyBytes, &tokenResp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("unable to read kyber response: %v", err)))
			return
		}

		// Send retrieved access token to Kyber to access user information
		userResp, err := http.Get("https://kyber.network/api/user_info?access_token=" + tokenResp.AccessToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("unable to retrieve user info: %v", err)))
			return
		}
		defer userResp.Body.Close()

		userBytes, err := ioutil.ReadAll(userResp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("unable to read kyber response: %v", err)))
			return
		}

		var userData userResponse
		if err := json.Unmarshal(userBytes, &userData); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("unable to read kyber response: %v", err)))
			return
		}

		if userData.Status != statusApproved {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("trader is not authorized: kyber status = %v", userData.Status)))
			return
		}

		// Update verification time in database
		if err := loginAdapter.PostVerification(data.Address, int64(userData.UID), ingress.KYCKyber); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("failed to post verification: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(userBytes)
	}
}

// ApproveWithdrawalHandler handles all HTTP open order requests
func ApproveWithdrawalHandler(approveWithdrawalAdapter ApproveWithdrawalAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		approveWithdrawalRequest := ApproveWithdrawalRequest{}
		if err := json.NewDecoder(r.Body).Decode(&approveWithdrawalRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode json into approve withdrawal request: %v", err)))
			return
		}
		signature, err := approveWithdrawalAdapter.ApproveWithdrawal(approveWithdrawalRequest.Trader, approveWithdrawalRequest.TokenID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("failed to approve withdrawal: %v", err)))
			return
		}

		response, err := json.Marshal(ApproveWithdrawalResponse{
			Signature: MarshalSignature(signature),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("failed to marshal ApproveWithdrawalResponse: %v", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

// GetAddressHandler handles all HTTP get address requests
func GetAddressHandler(getAddressAdapter GetAddressAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		addr, err := getAddressAdapter.GetAddress(params["orderID"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("failed to find the required address: %v", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(addr))
	}
}

// PostAddressHandler handles all HTTP post address requests
func PostAddressHandler(postAddressAdapter PostAddressAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postAddressRequest := PostAddressRequest{}
		if err := json.NewDecoder(r.Body).Decode(&postAddressRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode json into post address request: %v", err)))
			return
		}

		if err := postAddressAdapter.PostAddress(postAddressRequest.Info, postAddressRequest.Signature); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("failed to post address: %v", err)))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// GetAuthorizedHandler handles all HTTP get authorized requests
func GetAuthorizedHandler(getAuthorizeAdapter GetAuthorizeAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		res := GetAuthorizeResponse{}
		addr, err := getAuthorizeAdapter.GetAuthorizedAddress(params["address"])
		if err != nil {
			if err == sql.ErrNoRows {
				res.Status = false
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("cannot get authorization status: %v", err)))
				return
			}
		} else {
			res.AtomAddress = addr
			res.Status = true
		}
		respBytes, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot encode json into the expected response format: %v", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	}
}

// PostAuthorizeHandler handles all HTTP post authorize requests
func PostAuthorizeHandler(postAuthorizeAdapter PostAuthorizeAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postAuthorizeRequest := PostAuthorizeRequest{}
		if err := json.NewDecoder(r.Body).Decode(&postAuthorizeRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode json into address and atom address: %v", err)))
			return
		}
		if err := postAuthorizeAdapter.PostAuthorizedAddress(postAuthorizeRequest.AtomAddress, postAuthorizeRequest.Signature); err != nil {
			if err == ErrUnauthorized {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(fmt.Sprintf("Signing address is not KYC'd: %v", err)))
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Failed to authorize: %v", err)))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// GetSwapHandler handles all HTTP get swap details requests
func GetSwapHandler(getSwapAdapter GetSwapAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		swap, err := getSwapAdapter.GetSwap(params["orderID"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("required swap details not found: %v", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(swap))
	}
}

// PostSwapHandler handles all HTTP get swap details requests
func PostSwapHandler(postSwapAdapter PostSwapAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postSwapRequest := PostSwapRequest{}
		if err := json.NewDecoder(r.Body).Decode(&postSwapRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode json into post swap request: %v", err)))
			return
		}

		if err := postSwapAdapter.PostSwap(postSwapRequest.Info, postSwapRequest.Signature); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("failed to save the swap datails: %v", err)))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetKYCHandler(ingressAdapter IngressAdapter, kyberID, kyberSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		address := params["address"]
		kycType, err := traderVerified(ingressAdapter, kyberID, kyberSecret, address)
		if err != nil {
			errString := fmt.Sprintf("cannot check trader verification: %v", err)
			log.Println(errString)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errString))
			raven.CaptureErrorAndWait(errors.New(errString), map[string]string{
				"trader": address,
			})
			return
		}
		if kycType == ingress.KYCNone {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("trader is not verified"))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func PostSwapCallbackHandler(ingressAdapter IngressAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var swap foundation.SwapBlob
		if err := json.NewDecoder(r.Body).Decode(&swap); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var info delayInfo
		if err := json.Unmarshal(swap.DelayInfo, &info); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		details, err := ingressAdapter.GetSwap(info.OrderID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		var pswap ingress.PartialSwap
		if err := json.Unmarshal([]byte(details), &pswap); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		swap.Delay = false
		swap.SendTo = pswap.SendTo
		swap.ReceiveFrom = pswap.ReceiveFrom
		swap.SendAmount = pswap.SendAmount
		swap.ReceiveAmount = pswap.ReceiveAmount
		data, err := json.Marshal(swap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// RecoveryHandler handles errors while processing the requests and populates the errors in the response
func RecoveryHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("%v", r)))
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func traderVerified(loginAdapter LoginAdapter, kyberID, kyberSecret, address string) (int, error) {
	disableKYC := os.Getenv("DISABLE_KYC") == "1"
	if disableKYC {
		return ingress.KYCWyre, nil
	}
	if address[:2] != "0x" {
		address = "0x" + address
	}
	verified, err := loginAdapter.WyreVerified(address)
	if err != nil {
		return ingress.KYCNone, fmt.Errorf("cannot check wyre verification: %v", err)
	}
	if verified {
		if err := loginAdapter.PostVerification(address, 0, ingress.KYCWyre); err != nil {
			return ingress.KYCNone, fmt.Errorf("cannot update wyre verification information in database: %v", err)
		}
		return ingress.KYCWyre, nil
	}

	// If the Wyre verification is unsuccessful, check if the
	// trader has verified using Kyber
	kyberUID, timestamp, err := loginAdapter.GetLogin(address)
	if err != nil {
		if err == sql.ErrNoRows {
			return ingress.KYCNone, nil
		}
		return ingress.KYCNone, fmt.Errorf("cannot get verification information from database: %v", err)
	}

	// Check to see if the trader has verified using Kyber in the last 24
	// hours
	if timestamp != "" {
		unix, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			return ingress.KYCNone, fmt.Errorf("cannot parse timestamp: %v", err)
		}
		if time.Unix(unix, 0).After(time.Now().AddDate(0, 0, -1)) {
			return ingress.KYCKyber, nil
		}
	}

	// If user has not verified recently, retrieve access token for interacting
	// with Kyber API
	urlString := "https://kyber.network/oauth/token"
	resp, err := http.PostForm(urlString, url.Values{"grant_type": {"client_credentials"}, "client_id": {kyberID}, "client_secret": {kyberSecret}})
	if err != nil {
		return ingress.KYCNone, fmt.Errorf("cannot send information to kyber: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ingress.KYCNone, fmt.Errorf("cannot read information from kyber: %v", err)
	}

	var tokenResp tokenResponse
	if err := json.Unmarshal(bodyBytes, &tokenResp); err != nil {
		return ingress.KYCNone, fmt.Errorf("cannot unmarshal kyber access token data: %v", err)
	}

	// Retrieve information for trader with uID
	resp, err = http.Get("https://kyber.network/api/authorized_users?access_token=" + tokenResp.AccessToken + "&uid=" + fmt.Sprintf("%v", kyberUID))
	if err != nil {
		return ingress.KYCNone, fmt.Errorf("cannot send user information to kyber: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ingress.KYCNone, fmt.Errorf("cannot read user information from kyber: %v", err)
	}

	var usersResp usersResponse
	if err := json.Unmarshal(bodyBytes, &usersResp); err != nil {
		return ingress.KYCNone, fmt.Errorf("cannot unmarshal authorized kyber users: %v", err)
	}

	// Submit verification if the selected address is still verified with the
	// trader's Kyber account
	if len(usersResp.Users) > 0 {
		for _, addr := range usersResp.Users[0].Addresses {
			if addr == address {
				if err := loginAdapter.PostVerification(address, kyberUID, ingress.KYCKyber); err != nil {
					return ingress.KYCNone, fmt.Errorf("cannot update kyber verification information in database: %v", err)
				}
				return ingress.KYCKyber, nil
			}
		}
		return ingress.KYCNone, nil
	}

	return ingress.KYCNone, nil
}

func traderApproved(address string, approvedTraders []string) bool {
	if address[:2] == "0x" {
		address = address[2:]
	}
	address = strings.ToLower(address)
	for _, trader := range approvedTraders {
		if strings.ToLower(trader) == address {
			return true
		}
	}
	return false
}

func rateLimit(limiter *rate.Limiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// netAddr := r.RemoteAddr
		// ipAddr := strings.Split(netAddr, ":")[0]
		if limiter.Allow() {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("too many request"))
	}
}
