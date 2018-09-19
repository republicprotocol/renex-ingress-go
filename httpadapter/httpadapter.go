package httpadapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/time/rate"
)

// NewIngressServer returns an http server that forwards requests to an
// IngressAdapter.
func NewIngressServer(ingressAdapter IngressAdapter) http.Handler {
	limiter := rate.NewLimiter(2, 5)
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/orders", rateLimit(limiter, OpenOrderHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/withdrawals", rateLimit(limiter, ApproveWithdrawalHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/address", rateLimit(limiter, PostAddressHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/swap", rateLimit(limiter, PostSwapHandler(ingressAdapter))).Methods("POST")
	r.HandleFunc("/authorize", rateLimit(limiter, PostAuthorizeHandler(ingressAdapter))).Methods("POST")
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
func OpenOrderHandler(openOrderAdapter OpenOrderAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		openOrderRequest := OpenOrderRequest{}
		if err := json.NewDecoder(r.Body).Decode(&openOrderRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode json into an order or a list of order fragments: %v", err)))
			return
		}
		signature, err := openOrderAdapter.OpenOrder(openOrderRequest.Address, openOrderRequest.OrderFragmentMappings)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot open order: %v", err)))
			return
		}

		response, err := json.Marshal(OpenOrderResponse{
			Signature: MarshalSignature(signature),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot open order: %v", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

// ApproveWithdrawalHandler handles all HTTP open order requests
func ApproveWithdrawalHandler(approveWithdrawalAdapter ApproveWithdrawalAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		approveWithdrawalRequest := ApproveWithdrawalRequest{}
		if err := json.NewDecoder(r.Body).Decode(&approveWithdrawalRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode json into a trader and token: %v", err)))
			return
		}
		signature, err := approveWithdrawalAdapter.ApproveWithdrawal(approveWithdrawalRequest.Trader, approveWithdrawalRequest.TokenID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot open order: %v", err)))
			return
		}

		response, err := json.Marshal(ApproveWithdrawalResponse{
			Signature: MarshalSignature(signature),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot open order: %v", err)))
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
			w.Write([]byte(fmt.Sprintf("cannot open order: %v", err)))
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
			w.Write([]byte(fmt.Sprintf("cannot decode json into a trader and token: %v", err)))
			return
		}

		if err := postAddressAdapter.PostAddress(postAddressRequest.OrderID, postAddressRequest.Address); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot open order: %v", err)))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// PostAuthorizeHandler handles all HTTP post Authorize requests
func PostAuthorizeHandler(postAuthorizeAdapter PostAuthorizeAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postAuthorizeRequest := PostAuthorizeRequest{}
		if err := json.NewDecoder(r.Body).Decode(&postAuthorizeRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot decode json into a trader and token: %v", err)))
			return
		}
		if err := postAuthorizeAdapter.PostAuthorize(postAuthorizeRequest.atomAddresses, postAuthorizeRequest.signature); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot open order: %v", err)))
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
			w.Write([]byte(fmt.Sprintf("cannot decode json into a trader and token: %v", err)))
			return
		}

		if err := postSwapAdapter.PostSwap(postSwapRequest.OrderID, postSwapRequest.Swap); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("cannot open order: %v", err)))
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

func rateLimit(limiter *rate.Limiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// netAddr := r.RemoteAddr
		// ipAddr := strings.Split(netAddr, ":")[0]
		if limiter.Allow() {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("too many requests"))
	}
}

func toBytes32(b []byte) ([32]byte, error) {
	bytes32 := [32]byte{}
	if len(b) != 32 {
		return bytes32, errors.New("length mismatch")
	}
	copy(bytes32[:], b[:32])
	return bytes32, nil
}
