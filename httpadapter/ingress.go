package httpadapter

import (
	"errors"

	"github.com/republicprotocol/renex-ingress-go/ingress"
)

// ErrInvalidSignatureLength is returned when a signature does not have the
// required length of 65 bytes.
var ErrInvalidSignatureLength = errors.New("invalid signature length")

// ErrInvalidOrderID is returned when an order ID is not consistent across
// OrderFragmentMappings in the same request.
var ErrInvalidOrderID = errors.New("invalid order id")

// ErrInvalidOrderIDLength is returned when an order ID does not have the
// required length of 32 bytes.
var ErrInvalidOrderIDLength = errors.New("invalid order id length")

// ErrInvalidOrderFragmentIDLength is returned when an order fragment ID does
// not have the required length of 32 bytes.
var ErrInvalidOrderFragmentIDLength = errors.New("invalid order fragment id length")

// ErrInvalidEncryptedCoExpShareLength is returned when an encrypted co-exp
// share does not contain exactly 2 encrypted values, an encrypted co and an
// encrypted exp.
var ErrInvalidEncryptedCoExpShareLength = errors.New("invalid encrypted co-exp share length")

// ErrInvalidPodHashLength is returned when a pod hash does not have the
// required length of 32 bytes.
var ErrInvalidPodHashLength = errors.New("invalid pod hash length")

// ErrEmptyOrderFragmentMapping is returned when an OrderFragmentMapping does
// not store any OrderFragments.
var ErrEmptyOrderFragmentMapping = errors.New("empty order fragment mapping")

// An OpenOrderAdapter can be used to open an order.Order by sending an
// OrderFragmentMapping to the Darknodes in the network.
type OpenOrderAdapter interface {
	OpenOrder(signature string, orderFragmentMappings OrderFragmentMappings) error
}

// A CancelOrderAdapter can be used to cancel an order.Order by sending a
// signed cancelation message to the Ethereum blockchain where all Darknodes in
// the network will observe it.
type CancelOrderAdapter interface {
	CancelOrder(signature string, orderID string) error
}

// An IngressAdapter implements the OpenOrderAdapter and the
// CancelOrderAdapter.
type IngressAdapter interface {
	OpenOrderAdapter
	CancelOrderAdapter
}

type ingressAdapter struct {
	ingress.Ingress
}

// NewIngressAdapter returns an IngressAdapter that marshals and unmarshals
// requests before forwarding the request to an Ingress service.
func NewIngressAdapter(ingress ingress.Ingress) IngressAdapter {
	return &ingressAdapter{
		Ingress: ingress,
	}
}

// OpenOrder implements the OpenOrderAdapter interface.
func (adapter *ingressAdapter) OpenOrder(signatureIn string, orderFragmentMappingsIn OrderFragmentMappings) error {
	signature, err := UnmarshalSignature(signatureIn)
	if err != nil {
		return err
	}

	orderID, orderFragmentMappings, err := UnmarshalOrderFragmentMappings(orderFragmentMappingsIn)
	if err != nil {
		return err
	}

	return adapter.Ingress.OpenOrder(
		signature,
		orderID,
		orderFragmentMappings,
	)
}

// CancelOrder implements the CancelOrderAdapter interface.
func (adapter *ingressAdapter) CancelOrder(signatureIn string, orderIDIn string) error {
	signature, err := UnmarshalSignature(signatureIn)
	if err != nil {
		return err
	}

	orderID, err := UnmarshalOrderID(orderIDIn)
	if err != nil {
		return err
	}

	return adapter.Ingress.CancelOrder(signature, orderID)
}
