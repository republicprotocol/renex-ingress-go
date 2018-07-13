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
	r.HandleFunc("/orders", OpenOrderHandler(openOrderAdapter)).Methods("POST")
	r.HandleFunc("/orders", CancelOrderHandler(cancelOrderAdapter)).Methods("DELETE")
	r.Use(http.RecoveryHandler)

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
		if err := openOrderAdapter.OpenOrder(openOrderRequest.Signature, openOrderRequest.OrderFragmentMappings); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot open order: %v", err)))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// CancelOrderHandler handles HTTP Delete Requests
func CancelOrderHandler(cancelOrderAdapter CancelOrderAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID := r.URL.Query().Get("id")
		if orderID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot cancel order: nil id")))
			return
		}
		signature := r.URL.Query().Get("signature")
		if signature == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot cancel order: nil signature")))
			return
		}
		if err := cancelOrderAdapter.CancelOrder(signature, orderID); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("cannot cancel order: %v", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
