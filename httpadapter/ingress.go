package httpadapter

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/republicprotocol/renex-ingress-go/ingress"
	renex "github.com/republicprotocol/renex-sdk-go"
	"github.com/republicprotocol/republic-go/order"
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

type GetRewardsAdapter interface {
	GetRewards(string) (map[string]*big.Int, error)
}

type KYCAdapter interface {
}

type LoginAdapter interface {
	GetLogin(address string) (int64, string, error)
	PostLogin(address, referrer string) error
	PostVerification(address string, kyberUID int64, kycType int) error
	WyreVerified(traderIn string) (bool, error)
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
	GetRewardsAdapter
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

func (adapter *ingressAdapter) GetRewards(address string) (map[string]*big.Int, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	ren, err := renex.NewRenExWithPrivKey("mainnet", privKey)
	if err != nil {
		return nil, err
	}

	// Calculate amount trader should receive for orders they have opened
	rewards, err := adapter.getRewards(ren, address, big.NewInt(20)) // TODO: Confirm divisor values
	if err != nil {
		return nil, err
	}
	// TODO: Only get rewards after timestamp?

	// Calculate amount trader should receive for orders opened by referred
	// traders
	referrents, err := adapter.SelectReferrents(address)
	if err != nil {
		return nil, err
	}
	for _, referrent := range referrents {
		referralRewards, err := adapter.getRewards(ren, referrent, big.NewInt(5))
		if err != nil {
			return nil, err
		}
		for key, value := range referralRewards {
			rewards[key] = new(big.Int).Add(rewards[key], value)
		}
	}

	return rewards, nil
}

func (adapter *ingressAdapter) getRewards(ren renex.RenEx, address string, divisor *big.Int) (map[string]*big.Int, error) {
	rewards := make(map[string]*big.Int)
	orders, err := ren.Orderbook.ListOrdersByTrader(address)
	if err != nil {
		return nil, err
	}
	for _, orderID := range orders {
		settled, err := ren.Settled(orderID)
		if err != nil {
			return nil, err
		}
		if !settled {
			continue
		}
		orderMatch, err := ren.MatchDetails(orderID)
		if err != nil {
			return nil, err
		}
		var token string
		var newReward *big.Int
		if orderMatch.OrderIsBuy {
			token = order.Token(orderMatch.PriorityToken).String()
			newReward = new(big.Int).Div(orderMatch.PriorityFee, divisor)
		} else {
			token = order.Token(orderMatch.SecondaryToken).String()
			newReward = new(big.Int).Div(orderMatch.SecondaryFee, divisor)
		}
		rewards[token] = new(big.Int).Add(rewards[token], newReward)
	}
	return rewards, nil
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
