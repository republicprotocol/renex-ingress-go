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

var ErrUnauthorized = errors.New("unauthorized address")

// An OpenOrderAdapter can be used to open an order.Order by sending an
// OrderFragmentMapping to the Darknodes in the network.
type OpenOrderAdapter interface {
	OpenOrder(traderIn string, orderFragmentMappings OrderFragmentMappings) ([65]byte, error)
}

type ApproveWithdrawalAdapter interface {
	ApproveWithdrawal(traderIn string, tokenID uint32) ([65]byte, error)
}

type LoginAdapter interface {
	GetLogin(address string) (int64, string, error)
	PostLogin(address, referrer string) error
	PostVerification(address string, kyberUID int64, kycType int) error
	WyreVerified(traderIn string) (bool, error)
}

type OrderAdapter interface {
	InsertPartialSwap(swap ingress.PartialSwap) error
	PartialSwap(id string) (ingress.PartialSwap, error)
	FinalizedSwap(id string) (ingress.FinalizedSwap, bool, error)
}

// An IngressAdapter implements the OpenOrderAdapter and the
// ApproveWithdrawalAdapter.
type IngressAdapter interface {
	OpenOrderAdapter
	ApproveWithdrawalAdapter
	LoginAdapter
	OrderAdapter
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

func (adapter *ingressAdapter) WyreVerified(traderIn string) (bool, error) {
	trader, err := UnmarshalAddress(traderIn)
	if err != nil {
		return false, err
	}

	return adapter.Ingress.WyreVerified(trader)
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

func (adapter *ingressAdapter) GetLogin(address string) (int64, string, error) {
	return adapter.SelectLogin(address)
}

func (adapter *ingressAdapter) PostLogin(address, referrer string) error {
	return adapter.InsertLogin(address, referrer)
}

func (adapter *ingressAdapter) PostVerification(address string, kyberUID int64, kycType int) error {
	return adapter.UpdateLogin(address, kyberUID, kycType)
}

func (adapter *ingressAdapter) InsertPartialSwap(swap ingress.PartialSwap) error {
	return adapter.Ingress.InsertPartialSwap(swap)
}

func (adapter *ingressAdapter) PartialSwap(id string) (ingress.PartialSwap, error) {
	return adapter.Ingress.PartialSwap(id)
}

func (adapter *ingressAdapter) FinalizedSwap(id string) (ingress.FinalizedSwap, bool, error) {
	return adapter.Ingress.FinalizedSwap(id)
}
