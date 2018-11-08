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
	"strings"

	raven "github.com/getsentry/raven-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/time/rate"
)

type loginRequest struct {
	Address  string `json:"address"`
	Referrer string `json:"referrer"`
}

type verificationRequest struct {
	Address string `json:"address"`
	UID     string `json:"uid"`
	KYCType int    `json:"kycType"`
}

// Kyber request and response types
type authRequest struct {
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
	Status    string   `json:"kyc_status"`
	Addresses []string `json:"active_wallets"`
}

const (
	statusApproved string = "approved"
	statusPending  string = "pending"
	statusNone     string = "none"
)

// NewIngressServer returns an http server that forwards requests to an
// IngressAdapter.

func NewIngressServer(ingressAdapter IngressAdapter, approvedTraders []string, kyberSecret string) http.Handler {
	limiter := rate.NewLimiter(3, 20)
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/orders", rateLimit(limiter, OpenOrderHandler(ingressAdapter, approvedTraders))).Methods("POST")
	r.HandleFunc("/kyber", rateLimit(limiter, KyberKYCHandler(ingressAdapter, kyberSecret))).Methods("POST")
	r.HandleFunc("/login", rateLimit(limiter, LoginHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/verify", rateLimit(limiter, VerificationHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/withdrawals", rateLimit(limiter, ApproveWithdrawalHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/address", rateLimit(limiter, PostAddressHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/swap", rateLimit(limiter, PostSwapHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/authorize", rateLimit(limiter, PostAuthorizeHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/authorized/{address}", rateLimit(limiter, GetAuthorizedHandler(ingressAdapter))).Methods("GET")
	r.HandleFunc("/address/{orderID}", rateLimit(limiter, GetAddressHandler(ingressAdapter))).Methods("GET")
	r.HandleFunc("/swap/{orderID}", rateLimit(limiter, GetSwapHandler(ingressAdapter))).Methods("GET")
	r.Use(RecoveryHandler)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)

	return handler
}

// OpenOrderHandler handles all HTTP open order requests
func OpenOrderHandler(ingressAdapter IngressAdapter, approvedTraders []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		openOrderRequest := OpenOrderRequest{}
		if err := json.NewDecoder(r.Body).Decode(&openOrderRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode json into an order or a list of order fragments: %v", err)))
			return
		}
		// First check if trader has been manually approved (e.g. Lotan traders)
		if !traderApproved(openOrderRequest.Address, approvedTraders) {
			verified, err := traderVerified(ingressAdapter, openOrderRequest.Address)
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
			if !verified {
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

// KyberKYCHandler handles all Kyber KYC verification requests
func KyberKYCHandler(kycAdapter KYCAdapter, kyberSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode POST request data
		decoder := json.NewDecoder(r.Body)
		var data authRequest
		err := decoder.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode data: %v", err)))
			return
		}

		// Construct new request object with Kyber secret key
		data.Secret = kyberSecret
		byteArray, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot marshal data: %v", err)))
			return
		}

		// Forward updated request data to Kyber
		url := "https://kyber.network/oauth/token"
		postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(byteArray))
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

		w.WriteHeader(http.StatusOK)
		w.Write(userBytes)
	}
}

// LoginHandler handles trader address requests prior to KYC
func LoginHandler(loginAdapter LoginAdapter) http.HandlerFunc {
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
		_, err = loginAdapter.GetLogin(data.Address)
		if err != nil {
			if err == sql.ErrNoRows {
				if err := loginAdapter.PostLogin(data.Address, data.Referrer); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(fmt.Sprintf("cannot store login address: %v", err)))
					return
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("cannot retrieve login information: %v", err)))
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}

// VerificationHandler handles trader KYC verification requests
func VerificationHandler(verificationAdapter VerificationAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode POST request data
		decoder := json.NewDecoder(r.Body)
		var data verificationRequest
		err := decoder.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode data: %v", err)))
			return
		}

		verified, err := traderVerified(verificationAdapter, data.Address)
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
		if !verified {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("trader is not verified"))
			return
		}

		// Store verification information in database
		if err := verificationAdapter.PostVerification(data.Address, data.UID, data.KYCType); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot store verification data: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
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
			w.Write([]byte(fmt.Sprintf("jailed to marshal ApproveWithdrawalResponse: %v", err)))
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

func traderVerified(verificationAdapter VerificationAdapter, address string) (bool, error) {
	if !strings.HasPrefix(address, "0x") {
		address = "0x" + address
	}
	verified, err := verificationAdapter.WyreVerified(address)
	if err != nil {
		return false, err
	}
	if verified {
		return true, nil
	}

	// If the Wyre verification is unsuccessful, check if the
	// trader has verified using Kyber.
	_, err = verificationAdapter.GetTrader(address)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
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

func toBytes32(b []byte) ([32]byte, error) {
	bytes32 := [32]byte{}
	if len(b) != 32 {
		return bytes32, errors.New("Length mismatch")
	}
	copy(bytes32[:], b[:32])
	return bytes32, nil
}
