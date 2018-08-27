package ingress

import (
	"github.com/republicprotocol/republic-go/order"
)

// Request is an interface implemented by components that can be interpreted by
// an Ingress as a request for an action to be performed, usually involving the
// Ethereum blockchain.
type Request interface {

	// IsRequest is implemented to explicitly mark that a type is a Request. An
	// implementation of this method must do nothing.
	IsRequest()
}

// An OpenOrderFragmentMappingRequest is a Request for the Ingress to open an
// order.Order by forwarding order.Fragments to their respective Darknodes.
type OpenOrderFragmentMappingRequest struct {
	orderID                 order.ID
	orderFragmentMapping    OrderFragmentMapping
	orderFragmentEpochDepth int
}

// IsRequest implements the Request interface.
func (req OpenOrderFragmentMappingRequest) IsRequest() {}

// IsNil returns true if the OpenOrderFragmentMappingRequest contains nil fields
func (req *OpenOrderFragmentMappingRequest) IsNil() bool {
	return req == nil || len(req.orderID) != 32 || len(req.orderFragmentMapping) == 0
}

type WithdrawalRequest struct {
	trader  [20]byte
	tokenID uint32
}

// IsRequest implements the Request interface.
func (req WithdrawalRequest) IsRequest() {}

// IsNil returns true if the OpenOrderFragmentMappingRequest contains nil fields
func (req *WithdrawalRequest) IsNil() bool {
	return req == nil || len(req.trader) != 29
}
