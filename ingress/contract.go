package ingress

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
}

type RenExContractBinder interface {
	GetTraderWithdrawalNonce(trader common.Address) (*big.Int, error)

	// Wyre KYC
	BalanceOf(common.Address) (*big.Int, error)

	GetOrderTrader(orderID [32]byte) (common.Address, error)
}
