package http

import (
	"encoding/json"
	"fmt"
	netHttp "net/http"

	"github.com/gorilla/mux"
	"github.com/republicprotocol/renex-ingress-api-go/httpadapter"
	"github.com/republicprotocol/republic-go/http"
	"github.com/rs/cors"
)

// OpenOrderRequest is an JSON object sent to the HTTP handlers to request the
// opening of an order.
type OpenOrderRequest struct {
	Signature             string                            `json:"signature"`
	OrderFragmentMappings httpadapter.OrderFragmentMappings `json:"orderFragmentMappings"`
}

func NewServer(openOrderAdapter httpadapter.OpenOrderAdapter, cancelOrderAdapter httpadapter.CancelOrderAdapter) netHttp.Handler {
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
func OpenOrderHandler(adapter httpadapter.OpenOrderAdapter) netHttp.HandlerFunc {
	return func(w netHttp.ResponseWriter, r *netHttp.Request) {
		openOrderRequest := OpenOrderRequest{}
		if err := json.NewDecoder(r.Body).Decode(&openOrderRequest); err != nil {
			http.WriteError(w, netHttp.StatusBadRequest, fmt.Sprintf("cannot decode json into an order or a list of order fragments: %v", err))
			return
		}
		if err := adapter.OpenOrder(openOrderRequest.Signature, openOrderRequest.OrderFragmentMappings); err != nil {
			http.WriteError(w, netHttp.StatusInternalServerError, fmt.Sprintf("cannot open order: %v", err))
			return
		}
		w.WriteHeader(netHttp.StatusCreated)
	}
}

// CancelOrderHandler handles HTTP Delete Requests
func CancelOrderHandler(adapter httpadapter.CancelOrderAdapter) netHttp.HandlerFunc {
	return func(w netHttp.ResponseWriter, r *netHttp.Request) {
		orderID := r.URL.Query().Get("id")
		if orderID == "" {
			http.WriteError(w, netHttp.StatusBadRequest, fmt.Sprintf("cannot cancel order: nil id"))
			return
		}
		signature := r.URL.Query().Get("signature")
		if signature == "" {
			http.WriteError(w, netHttp.StatusBadRequest, fmt.Sprintf("cannot cancel order: nil signature"))
			return
		}
		if err := adapter.CancelOrder(signature, orderID); err != nil {
			http.WriteError(w, netHttp.StatusInternalServerError, fmt.Sprintf("cannot cancel order: %v", err))
			return
		}
		w.WriteHeader(netHttp.StatusOK)
	}
}
