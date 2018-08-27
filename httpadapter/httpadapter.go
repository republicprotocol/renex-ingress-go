package httpadapter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// NewIngressServer returns an http server that forwards requests to an
// IngressAdapter.
func NewIngressServer(ingressAdapter IngressAdapter) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/orders", OpenOrderHandler(ingressAdapter)).Methods("POST")
	r.HandleFunc("/withdrawals", ApproveWithdrawalHandler(ingressAdapter)).Methods("POST")
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
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(MarshalSignature(signature)))
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
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(MarshalSignature(signature)))
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
