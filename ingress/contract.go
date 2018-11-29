package ingress

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/republicprotocol/renex-ingress-go/contract"
	"github.com/republicprotocol/republic-go/order"
	"github.com/republicprotocol/republic-go/registry"
)

// ContractBinder will define all methods that the ingress will require to
// communicate with smart contracts.
type ContractBinder interface {
	MinimumEpochInterval() (*big.Int, error)

	Epoch() (registry.Epoch, error)

	NextEpoch() (registry.Epoch, error)

	PreviousEpoch() (registry.Epoch, error)

	Pods() ([]registry.Pod, error)

	PreviousPods() ([]registry.Pod, error)

	GetOrderTrader(orderID [32]byte) (common.Address, error)

	GetMatchDetails(orderID order.ID) (contract.OrderMatch, error)

	Orders(offset, limit int) ([]order.ID, []order.Status, []string, error)

	OrderCounts() (uint64, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)
}

type RenExContractBinder interface {
	GetTraderWithdrawalNonce(trader common.Address) (*big.Int, error)

	// Wyre KYC
	BalanceOf(common.Address) (*big.Int, error)
}
