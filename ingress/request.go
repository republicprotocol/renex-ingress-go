package ingress

import "github.com/republicprotocol/republic-go/order"

// Request is an interface implemented by components that can be interpreted by
// an Ingress as a request for an action to be performed, usually involving the
// Ethereum blockchain.
type Request interface {

	// IsRequest is implemented to explicitly mark that a type is a Request. An
	// implementation of this method must do nothing.
	IsRequest()
}

// An OpenOrderRequest is a Request for the Ingress to open an order.Order on
// the Ethereum blockchain.
type OpenOrderRequest struct {
	signature   [65]byte
	orderID     order.ID
	orderParity order.Parity
}

// IsRequest implements the Request interface.
func (req OpenOrderRequest) IsRequest() {}

// IsNil returns true if the OpenOrderRequest contains nil fields
func (req *OpenOrderRequest) IsNil() bool {
	return req == nil || len(req.signature) != 65 || len(req.orderID) != 32
}

// An OpenOrderFragmentMappingRequest is a Request for the Ingress to open an
// order.Order by forwarding order.Fragments to their respective Darknodes.
type OpenOrderFragmentMappingRequest struct {
	signature               [65]byte
	orderID                 order.ID
	orderFragmentMapping    OrderFragmentMapping
	orderFragmentEpochDepth int
}

// IsRequest implements the Request interface.
func (req OpenOrderFragmentMappingRequest) IsRequest() {}

// IsNil returns true if the OpenOrderFragmentMappingRequest contains nil fields
func (req *OpenOrderFragmentMappingRequest) IsNil() bool {
	return req == nil || len(req.signature) != 65 || len(req.orderID) != 32 || len(req.orderFragmentMapping) == 0
}

// A CancelOrderRequest is a Request for the Ingress to cancel an order.Order
// on the Ethereum blockchain.
type CancelOrderRequest struct {
	signature [65]byte
	orderID   order.ID
}

// IsRequest implements the Request interface.
func (req CancelOrderRequest) IsRequest() {}

// IsNil returns true if the CancelOrderRequest contains nil fields
func (req *CancelOrderRequest) IsNil() bool {
	return req == nil || len(req.signature) != 65 || len(req.orderID) != 32
}
