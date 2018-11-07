package httpadapter

import (
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	WyreVerified(traderIn string) (bool, error)
	GetTrader(address string) (string, error)
}

type ApproveWithdrawalAdapter interface {
	ApproveWithdrawal(traderIn string, tokenID uint32) ([65]byte, error)
}

type GetAddressAdapter interface {
	GetAddress(string) (string, error)
}

type PostAddressAdapter interface {
	PostAddress(PostAddressInfo, string) error
}

type GetSwapAdapter interface {
	GetSwap(string) (string, error)
}

type PostSwapAdapter interface {
	PostSwap(PostSwapInfo, string) error
}

type PostAuthorizeAdapter interface {
	PostAuthorizedAddress(string, string) error
}

type GetAuthorizeAdapter interface {
	GetAuthorizedAddress(string) (string, error)
}

type KYCAdapter interface {
	PostTrader(string) error
}

type LoginAdapter interface {
	GetLogin(address string) (string, error)
	PostLogin(address, referral string) error
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
	PostAuthorizeAdapter
	GetAuthorizeAdapter
	KYCAdapter
	LoginAdapter
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

func (adapter *ingressAdapter) GetAddress(orderID string) (string, error) {
	return adapter.SelectAddress(orderID)
}

func (adapter *ingressAdapter) PostAddress(info PostAddressInfo, signature string) error {
	infoBytes, err := json.Marshal(info)
	if err != nil {
		return err
	}
	hash := crypto.Keccak256(infoBytes)
	sigBytes, err := UnmarshalSignature(signature)
	if err != nil {
		return err
	}

	publicKey, err := crypto.SigToPub(hash, sigBytes[:])
	if err != nil {
		return err
	}
	address := crypto.PubkeyToAddress(*publicKey)
	if err := adapter.IsAuthorized(info.OrderID, address.String()); err != nil {
		return err
	}
	return adapter.InsertAddress(info.OrderID, info.Address)
}

func (adapter *ingressAdapter) GetSwap(orderID string) (string, error) {
	return adapter.SelectSwapDetails(orderID)
}

func (adapter *ingressAdapter) PostSwap(info PostSwapInfo, signature string) error {
	swapBytes, err := json.Marshal(info)
	if err != nil {
		return err
	}
	hash := crypto.Keccak256(swapBytes)
	sigBytes, err := UnmarshalSignature(signature)
	if err != nil {
		return err
	}

	publicKey, err := crypto.SigToPub(hash, sigBytes[:])
	if err != nil {
		return err
	}
	address := crypto.PubkeyToAddress(*publicKey)
	if err := adapter.IsAuthorized(info.OrderID, address.String()); err != nil {
		return err
	}
	return adapter.InsertSwapDetails(info.OrderID, info.Swap)
}

func (adapter *ingressAdapter) PostAuthorizedAddress(addr, signature string) error {
	address := common.HexToAddress(addr)

	message := append([]byte("RenEx: authorize: "), address.Bytes()...)
	signatureData := append([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))), message...)

	hash := crypto.Keccak256(signatureData)
	sigBytes, err := UnmarshalSignature(signature)
	if err != nil {
		return err
	}

	publicKey, err := crypto.SigToPub(hash, sigBytes[:])
	if err != nil {
		return err
	}

	kycAddress := crypto.PubkeyToAddress(*publicKey)

	// TODO: Removed for test purposes
	// verified, err := traderVerified(adapter, kycAddress.String())
	// if err != nil {
	// 	return err
	// }
	// if !verified {
	// 	return errors.New("address not verified")
	// }
	return adapter.InsertAuthorizedAddress(kycAddress.String(), addr)
}

func (adapter *ingressAdapter) GetAuthorizedAddress(addr string) (string, error) {
	return adapter.SelectAuthorizedAddress(addr)
}

func (adapter *ingressAdapter) GetTrader(address string) (string, error) {
	return adapter.SelectTrader(address)
}

func (adapter *ingressAdapter) PostTrader(address string) error {
	return adapter.InsertTrader(address)
}

func (adapter *ingressAdapter) GetLogin(address string) (string, error) {
	return adapter.SelectLogin(address)
}

func (adapter *ingressAdapter) PostLogin(address, referral string) error {
	return adapter.InsertLogin(address, referral)
}

func (adapter *ingressAdapter) IsAuthorized(orderID string, address string) error {
	hexBytes, err := hex.DecodeString(orderID)
	if err != nil {
		return err
	}
	id, err := UnmarshalOrderID(base64.StdEncoding.EncodeToString(hexBytes))
	if err != nil {
		return err
	}
	addr, err := adapter.Ingress.GetOrderTrader(id)
	if err != nil {
		return err
	}
	atomAddr, err := adapter.GetAuthorizedAddress(addr.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUnauthorized
		}
		return err
	}
	if address != atomAddr {
		return ErrUnauthorized
	}
	return nil
}
