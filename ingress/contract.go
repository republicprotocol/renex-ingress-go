package ingress

import (
	"math/big"

	"github.com/republicprotocol/republic-go/order"
	"github.com/republicprotocol/republic-go/registry"
)

// ContractBinder will define all methods that the ingress will require to
// communicate with smart contracts.
type ContractBinder interface {
	OpenBuyOrder(signature [65]byte, orderID order.ID) error

	OpenSellOrder(signature [65]byte, orderID order.ID) error

	CancelOrder(signature [65]byte, orderID order.ID) error

	MinimumEpochInterval() (*big.Int, error)

	Epoch() (registry.Epoch, error)

	NextEpoch() (registry.Epoch, error)

	PreviousEpoch() (registry.Epoch, error)

	Pods() ([]registry.Pod, error)

	PreviousPods() ([]registry.Pod, error)
}
