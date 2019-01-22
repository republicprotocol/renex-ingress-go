package httpadapter

import (
	"bytes"
	"database/sql"
	"encoding/base64"
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

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/getsentry/raven-go"
	"github.com/gorilla/mux"
	"github.com/republicprotocol/renex-ingress-go/ingress"
	"github.com/republicprotocol/swapperd/foundation/blockchain"
	"github.com/republicprotocol/swapperd/foundation/swap"
	"github.com/rs/cors"
	"golang.org/x/crypto/sha3"
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

type Message struct {
	KycAddr          string `json:"kycAddr"`
	OrderID          string `json:"orderID"`
	ReceiveTokenAddr string `json:"receiveTokenAddr"`
	SendTokenAddr    string `json:"sendTokenAddr"`
}

type delayInfo struct {
	Message   Message `json:"message"`
	Signature string  `json:"signature"`
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
	r.HandleFunc("/kyc/{address}", rateLimit(limiter, GetKYCHandler(ingressAdapter, kyberID, kyberSecret))).Methods("GET")
	r.HandleFunc("/orders", rateLimit(limiter, PostOrderHandler(ingressAdapter, approvedTraders, kyberID, kyberSecret))).Methods("POST")
	r.HandleFunc("/login", rateLimit(limiter, PostLoginHandler(ingressAdapter, kyberID, kyberSecret))).Methods("POST")
	r.HandleFunc("/kyber", rateLimit(limiter, PostKyberHandler(ingressAdapter, kyberID, kyberSecret))).Methods("POST")
	r.HandleFunc("/withdrawals", rateLimit(limiter, PostWithdrawalHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/swapperd/cb", rateLimit(limiter, PostSwapCallbackHandler(ingressAdapter, kyberID, kyberSecret))).Methods("POST")
	r.HandleFunc("/authorize", rateLimit(limiter, PostAuthorizeHandler(ingressAdapter, kyberID, kyberSecret))).Methods("POST")
	r.Use(RecoveryHandler)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)

	return handler
}

// PostOrderHandler handles all HTTP open order requests
func PostOrderHandler(ingressAdapter IngressAdapter, approvedTraders []string, kyberID, kyberSecret string) http.HandlerFunc {
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

// PostLoginHandler handles trader login requests
func PostLoginHandler(loginAdapter LoginAdapter, kyberID, kyberSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode POST request data
		var data loginRequest
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			errString := fmt.Sprintf("cannot decode data: %v", err)
			log.Println(errString)
			http.Error(w, errString, http.StatusBadRequest)
			return
		}

		// Store address in database if it does not already exist
		if err := loginAdapter.PostLogin(data.Address, data.Referrer); err != nil {
			errString := fmt.Sprintf("cannot store login address: %v", err)
			log.Println(errString)
			http.Error(w, errString, http.StatusInternalServerError)
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
			http.Error(w, errString, http.StatusInternalServerError)
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

// PostKyberHandler handles all Kyber authorization requests
func PostKyberHandler(loginAdapter LoginAdapter, kyberID, kyberSecret string) http.HandlerFunc {
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

// PostWithdrawalHandler handles all HTTP open order requests
func PostWithdrawalHandler(approveWithdrawalAdapter ApproveWithdrawalAdapter) http.HandlerFunc {
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

func PostSwapCallbackHandler(ingressAdapter IngressAdapter, kyberID, kyberSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var blob swap.SwapBlob
		if err := json.NewDecoder(r.Body).Decode(&blob); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var info delayInfo
		log.Println("callback for swapID:", blob.ID)
		if err := json.Unmarshal(blob.DelayInfo, &info); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		messageByte, err := json.Marshal(info.Message)
		if err != nil {
			http.Error(w, "unable to marshal the message", http.StatusBadRequest)
			return
		}

		// verify request
		hash := sha3.Sum256(messageByte)
		sigBytes, err := base64.StdEncoding.DecodeString(info.Signature)
		if err != nil {
			log.Println("unable marshal the signature", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		publicKey, err := crypto.SigToPub(hash[:], sigBytes)
		if err != nil {
			log.Println("unable verify signature address", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		signerAddr := crypto.PubkeyToAddress(*publicKey).Hex()
		kycType, err := traderVerified(ingressAdapter, kyberID, kyberSecret, signerAddr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if kycType == 0 {
			http.Error(w, "trader not kyced", http.StatusUnauthorized)
			return
		}

		// return the finalized blob if we have the finalized blob
		pSwap := ingress.PartialSwap{
			OrderID:     info.Message.OrderID,
			KycAddr:     info.Message.KycAddr,
			SendTo:      info.Message.SendTokenAddr,
			ReceiveFrom: info.Message.ReceiveTokenAddr,
			SecretHash:  blob.SecretHash,
			TimeLock:    time.Now().Add(48 * time.Hour).Unix(),
		}
		defer ingressAdapter.InsertPartialSwap(pSwap)

		// Check if we have the finalized blob info.
		finalizedSwap, canceled, err := ingressAdapter.FinalizedSwap(pSwap.OrderID)
		if err != nil {
			log.Println("no content:", err)
			http.Error(w, err.Error(), http.StatusNoContent)
			return
		}

		if canceled {
			http.Error(w, "order has been canceled", http.StatusGone)
			return
		}

		// Fill missing fields in the blob and return
		blob.Delay = false
		blob.SendTo = finalizedSwap.SendTo
		blob.ReceiveFrom = finalizedSwap.ReceiveFrom
		blob.SendAmount = finalizedSwap.SendAmount
		blob.ReceiveAmount = finalizedSwap.ReceiveAmount
		blob.ShouldInitiateFirst = finalizedSwap.ShouldInitiateFirst
		blob.TimeLock = finalizedSwap.TimeLock
		blob.SecretHash = finalizedSwap.SecretHash

		sendToken, err := blockchain.PatchToken(string(blob.SendToken))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		blob.BrokerSendTokenAddr, err = brokerAddress(sendToken.Blockchain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		receiveToken, err := blockchain.PatchToken(string(blob.ReceiveToken))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		blob.BrokerReceiveTokenAddr, err = brokerAddress(receiveToken.Blockchain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := json.Marshal(blob)
		if err != nil {
			log.Println("cannot marshal blob", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		log.Printf("%+v", blob)
	}
}

func PostAuthorizeHandler(ingressAdapter IngressAdapter, kyberID, kyberSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request
		var auth PostAuthorizeRequest
		if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
			handleErr(w, fmt.Sprintf("cannot decode request, %v", err), http.StatusBadRequest)
			return
		}

		// Extract the signer's address
		signatureData := append([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(common.Hex2Bytes(auth.Address)))), common.Hex2Bytes(auth.Address)...)
		hash := sha3.Sum256(signatureData)
		sigBytes, err := base64.StdEncoding.DecodeString(auth.Signature)
		if err != nil {
			handleErr(w, fmt.Sprintf("unable marshal the signature, %v", err), http.StatusInternalServerError)
			return
		}
		publicKey, err := crypto.SigToPub(hash[:], sigBytes)
		if err != nil {
			http.Error(w, fmt.Sprintf("unable verify signature address, %v", err), http.StatusInternalServerError)
			return
		}
		signerAddr := crypto.PubkeyToAddress(*publicKey).Hex()

		// Verify if the singer is kyced
		kycType, err := traderVerified(ingressAdapter, kyberID, kyberSecret, signerAddr)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v: signer %v", err.Error(), signerAddr), http.StatusUnauthorized)
			return
		}
		if kycType == 0 {
			http.Error(w, "trader not kyced", http.StatusUnauthorized)
			return
		}

		// todo : Handle the new address
		var address string
		switch len(auth.Address) {
		case 40:
			address = "0x" + auth.Address
		case 42:
			address = auth.Address
		default:
			handleErr(w, "invalid address", http.StatusBadRequest)
		}

		if err := ingressAdapter.Authorize(signerAddr, address); err != nil {
			handleErr(w, fmt.Sprintf("cannot store the new address, %v", err), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func brokerAddress(bcName blockchain.BlockchainName) (string, error) {
	switch bcName {
	case blockchain.Ethereum:
		return os.Getenv("ETH_VAULT"), nil
	case blockchain.Bitcoin:
		return os.Getenv("BTC_VAULT"), nil
	default:
		return "", blockchain.NewErrUnsupportedBlockchain(bcName)
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

func handleErr(w http.ResponseWriter, errMessage string, code int) {
	//
	log.Println(errMessage)
	http.Error(w, errMessage, code)
}
