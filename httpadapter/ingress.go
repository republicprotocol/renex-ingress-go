package httpadapter

import (
	"errors"

	"github.com/republicprotocol/renex-ingress-go/ingress"
)

// ErrInvalidSignatureLength is returned when a signature does not have the
// required length of 65 bytes.
var ErrInvalidSignatureLength = errors.New("invalid signature length")

// ErrInvalidAddressLength is returned when a address does not have the
// required length of 20 bytes.
var ErrInvalidAddressLength = errors.New("invalid address length")

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
	OpenOrder(traderIn string, orderFragmentMappings OrderFragmentMappings) ([65]byte, error)
	TraderVerified(traderIn string) (bool, error)
}

type ApproveWithdrawalAdapter interface {
	ApproveWithdrawal(traderIn string, tokenID uint32) ([65]byte, error)
}

type GetAddressAdapter interface {
	GetAddress(string) (string, error)
}

type PostAddressAdapter interface {
	PostAddress(string, string) error
}

type GetSwapAdapter interface {
	GetSwap(string) (string, error)
}

type PostSwapAdapter interface {
	PostSwap(string, string) error
}

// An IngressAdapter implements the OpenOrderAdapter and the
// ApproveWithdrawalAdapter.
type IngressAdapter interface {
	OpenOrderAdapter
	ApproveWithdrawalAdapter
	GetAddressAdapter
	PostAddressAdapter
	GetSwapAdapter
	PostSwapAdapter
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
func (adapter *ingressAdapter) OpenOrder(traderIn string, orderFragmentMappingsIn OrderFragmentMappings) ([65]byte, error) {
	trader, err := UnmarshalAddress(traderIn)
	if err != nil {
		return [65]byte{}, err
	}

	orderID, orderFragmentMappings, err := UnmarshalOrderFragmentMappings(orderFragmentMappingsIn)
	if err != nil {
		return [65]byte{}, err
	}

	return adapter.Ingress.OpenOrder(
		trader,
		orderID,
		orderFragmentMappings,
	)
}

func (adapter *ingressAdapter) TraderVerified(traderIn string) (bool, error) {
	trader, err := UnmarshalAddress(traderIn)
	if err != nil {
		return false, err
	}

	return adapter.Ingress.TraderVerified(trader)
}

// ApproveWithdrawal implements the ApproveWithdrawalAdapter interface.
func (adapter *ingressAdapter) ApproveWithdrawal(traderIn string, tokenIDIn uint32) ([65]byte, error) {
	trader, err := UnmarshalAddress(traderIn)
	if err != nil {
		return [65]byte{}, err
	}

	return adapter.Ingress.ApproveWithdrawal(
		trader,
		tokenIDIn,
	)
}

func (adapter *ingressAdapter) GetAddress(orderID string) (string, error) {
	return adapter.SelectAddress(orderID)
}

func (adapter *ingressAdapter) PostAddress(orderID string, address string) error {
	return adapter.InsertAddress(orderID, address)
}

func (adapter *ingressAdapter) GetSwap(orderID string) (string, error) {
	return adapter.SelectSwapDetails(orderID)
}

func (adapter *ingressAdapter) PostSwap(orderID string, swap string) error {
	return adapter.InsertSwapDetails(orderID, swap)
}
