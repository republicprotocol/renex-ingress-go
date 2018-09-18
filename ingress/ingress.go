package ingress

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/dispatch"
	"github.com/republicprotocol/republic-go/logger"
	"github.com/republicprotocol/republic-go/order"
	"github.com/republicprotocol/republic-go/orderbook"
	"github.com/republicprotocol/republic-go/registry"
	"github.com/republicprotocol/republic-go/swarm"
)

// ErrUnknownPod is returned when an unknown pod is mapped.
var ErrUnknownPod = errors.New("unknown pod id")

// ErrUnsupportedEpochDepth is returned when an unsupported epoch depth is
// received in an OrderFragmentMapping.
var ErrUnsupportedEpochDepth = errors.New("unsupported epoch depth")

// ErrInvalidNumberOfPods is returned when an insufficient number of pods are
// mapped.
var ErrInvalidNumberOfPods = errors.New("invalid number of pods")

// ErrInvalidNumberOfOrderFragments is returned when a pod is mapped to an
// insufficient number of order fragments, or too many order fragments.
var ErrInvalidNumberOfOrderFragments = errors.New("invalid number of order fragments")

// ErrInvalidOrderFragmentMapping is returned when an order fragment mapping is
// of an invalid length.
var ErrInvalidOrderFragmentMapping = errors.New("invalid order fragment mappings")

// ErrInvalidEpochDepth is returned when an invalid epoch depth is provided
// upon verification.
var ErrInvalidEpochDepth = errors.New("invalid epoch depth")

// ErrCannotOpenOrderFragments is returned when none of the pods were available
// to receive order fragments
var ErrCannotOpenOrderFragments = errors.New("cannot open order fragments: no pod received an order fragment")

// NumBackgroundWorkers is the number of background workers that the Ingress
// will use.
var NumBackgroundWorkers = runtime.NumCPU() * 4

// An OrderFragmentMapping maps pods to encrypted order fragments.
type OrderFragmentMapping map[[32]byte][]OrderFragment

// OrderFragmentMappings is a slice where the index of an OrderFragmentMapping
// represents the epoch depth of each OrderFragment inside the mapping.
type OrderFragmentMappings []OrderFragmentMapping

// OrderFragment has an order.EncryptedFragment, encrypted by the trader before
// being sent to the Ingress, and the required index that identifies which set
// shares are held by the order.EncryptedFragment.
type OrderFragment struct {
	order.EncryptedFragment
	Index int64
}

// Ingress interface can approve orders to opened on to the Orderbook and can
// forward fragments on
type Ingress interface {

	// Sync the epoch.
	Sync(<-chan struct{}) <-chan error

	// OpenOrder on the Darkpool and returns a signed approval for the order to
	// be opened in the Orderbook. The trader address and order ID are signed
	// together so that the approval is only valid for that trader. The order
	// fragment mapping is used to send order fragments to pods in the Darkpool.
	OpenOrder(trader [20]byte, orderID order.ID, orderFragmentMappings OrderFragmentMappings) ([65]byte, error)

	ApproveWithdrawal(trader [20]byte, tokenID uint32) ([65]byte, error)

	// ProcessRequests in the background. Closing the done channel will stop
	// all processing. Running this background worker is required to open and
	// cancel orders.
	ProcessRequests(done <-chan struct{}) <-chan error

	// Swapper interface implements atomic swapper network functions.
	Swapper
}

type ingress struct {
	ecdsaKey          crypto.EcdsaKey
	contract          ContractBinder
	renExContract     RenExContractBinder
	swarmer           swarm.Swarmer
	orderbookClient   orderbook.Client
	epochPollInterval time.Duration

	podsMu   *sync.RWMutex
	podsCurr map[[32]byte]registry.Pod
	podsPrev map[[32]byte]registry.Pod

	queueRequests chan Request
	Swapper
}

// NewIngress returns an Ingress. The background services of the Ingress must
// be started separately by calling Ingress.OpenOrderProcess and
// Ingress.OpenOrderFragmentsProcess.
func NewIngress(ecdsaKey crypto.EcdsaKey, contract ContractBinder, renExContract RenExContractBinder, swarmer swarm.Swarmer, orderbookClient orderbook.Client, epochPollInterval time.Duration, swapper Swapper) Ingress {
	ingress := &ingress{
		ecdsaKey:          ecdsaKey,
		contract:          contract,
		renExContract:     renExContract,
		swarmer:           swarmer,
		Swapper:           swapper,
		orderbookClient:   orderbookClient,
		epochPollInterval: epochPollInterval,

		podsMu:   new(sync.RWMutex),
		podsCurr: map[[32]byte]registry.Pod{},
		podsPrev: map[[32]byte]registry.Pod{},

		queueRequests: make(chan Request, 1024),
	}
	return ingress
}

// Sync implements the Ingress interface.
func (ingress *ingress) Sync(done <-chan struct{}) <-chan error {
	errs := make(chan error, 1)

	// Synchronize against the previous epoch
	epoch, err := ingress.contract.PreviousEpoch()
	if err != nil {
		errs <- err
		close(errs)
		return errs
	}

	pods, err := ingress.contract.PreviousPods()
	if err != nil {
		errs <- err
		close(errs)
		return errs
	}
	if err := ingress.syncFromEpoch(epoch, pods); err != nil {
		errs <- err
		close(errs)
		return errs
	}

	go func() {
		defer close(errs)

		dispatch.CoBegin(
			func() {
				ticker := time.NewTicker(ingress.epochPollInterval)
				defer ticker.Stop()

				for {
					select {
					case <-done:
						return
					case <-ticker.C:
					}

					// Get the current epoch
					nextEpoch, err := ingress.contract.Epoch()
					if err != nil {
						select {
						case <-done:
							return
						case errs <- err:
							continue
						}
					}

					// Check if it equals what we think the current epoch is
					// and update if necessary
					if bytes.Equal(epoch.Hash[:], nextEpoch.Hash[:]) {
						continue
					}
					epoch = nextEpoch
					pods, err := ingress.contract.Pods()
					if err != nil {
						select {
						case <-done:
							return
						case errs <- err:
							continue
						}
					}
					if err := ingress.syncFromEpoch(epoch, pods); err != nil {
						select {
						case <-done:
							return
						case errs <- err:
							continue
						}
					}
				}
			},
			func() {
				// ticker := time.NewTicker(2 * ingress.epochPollInterval)
				// defer ticker.Stop()
				//
				// for {
				// 	select {
				// 	case <-done:
				// 		return
				// 	case <-ticker.C:
				// 	}
				//
				// 	epoch, err := ingress.contract.NextEpoch()
				// 	if err != nil {
				// 		// Ignore the error to prevent verbose logging
				// 		continue
				// 	}
				// 	// Wait for a lower bound on the epoch
				// 	log.Printf("[info] (epoch) latest epoch = %v", base64.StdEncoding.EncodeToString(epoch.Hash[:]))
				// 	time.Sleep(time.Duration(epoch.BlockInterval.Int64()) * ingress.epochPollInterval)
				// }
			})
	}()

	return errs
}

func OpenOrderMessage(trader [20]byte, orderID order.ID) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, []byte("Republic Protocol: open: ")); err != nil {
		return []byte{}, err
	}
	if err := binary.Write(buf, binary.BigEndian, trader); err != nil {
		return []byte{}, err
	}
	if err := binary.Write(buf, binary.BigEndian, orderID); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func (ingress *ingress) OpenOrder(trader [20]byte, orderID order.ID, orderFragmentMappings OrderFragmentMappings) ([65]byte, error) {
	if err := ingress.verifyOrderFragmentMappings(orderFragmentMappings); err != nil {
		return [65]byte{}, err
	}

	log.Printf("[info] (open) signing order = %v", orderID)

	message, err := OpenOrderMessage(trader, orderID)
	if err != nil {
		return [65]byte{}, err
	}

	signatureData := append([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))), message...)
	fmt.Println("Trader:", hex.EncodeToString(trader[:]))
	fmt.Println("Order ID:", hex.EncodeToString(orderID[:]))
	fmt.Println("Signature data:", hex.EncodeToString(signatureData))
	hashedSignatureData := crypto.Keccak256(signatureData)
	fmt.Println("Hashed signature data:", hex.EncodeToString(hashedSignatureData))
	signature, err := ingress.ecdsaKey.Sign(hashedSignatureData)
	if err != nil {
		return [65]byte{}, err
	}
	fmt.Println("Signature:", hex.EncodeToString(signature))

	for i := range orderFragmentMappings {
		go func(i int) {
			log.Printf("[info] (open) queueing order fragments order = %v at depth = %v", orderID, i)
			ingress.queueRequests <- OpenOrderFragmentMappingRequest{
				orderID:                 orderID,
				orderFragmentMapping:    orderFragmentMappings[i],
				orderFragmentEpochDepth: i,
			}
		}(i)
	}

	var signature65 [65]byte
	copy(signature65[:], signature[:65])
	return signature65, nil
}

func WithdrawalMessage(trader [20]byte, tokenID uint32, traderNonce *big.Int) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, []byte("Republic Protocol: withdraw: ")); err != nil {
		return []byte{}, err
	}
	if err := binary.Write(buf, binary.BigEndian, trader); err != nil {
		return []byte{}, err
	}
	// NOTE: The contracts do not yet expect a tokenID - this may change to
	// restrict signatures to specific tokens
	// if err := binary.Write(buf, binary.BigEndian, tokenID); err != nil {
	// 	return []byte{}, err
	// }
	if err := binary.Write(buf, binary.BigEndian, uint64(0)); err != nil {
		return []byte{}, err
	}
	if err := binary.Write(buf, binary.BigEndian, uint64(0)); err != nil {
		return []byte{}, err
	}
	if err := binary.Write(buf, binary.BigEndian, uint64(0)); err != nil {
		return []byte{}, err
	}
	if err := binary.Write(buf, binary.BigEndian, traderNonce.Uint64()); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func (ingress *ingress) ApproveWithdrawal(trader [20]byte, tokenID uint32) ([65]byte, error) {
	log.Printf("[info] (open) approving withdrawal for %v", trader)
	// Append orderID

	// Retrieve trader nonce
	traderNonce, err := ingress.renExContract.GetTraderWithdrawalNonce(common.BytesToAddress(trader[:]))
	if err != nil {
		return [65]byte{}, err
	}

	message, err := WithdrawalMessage(trader, tokenID, traderNonce)
	if err != nil {
		return [65]byte{}, err
	}

	signatureData := append([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))), message...)
	hashedSignatureData := crypto.Keccak256(signatureData)
	signature, err := ingress.ecdsaKey.Sign(hashedSignatureData)
	if err != nil {
		return [65]byte{}, err
	}

	var signature65 [65]byte
	copy(signature65[:], signature[:65])
	return signature65, nil
}

func (ingress *ingress) ProcessRequests(done <-chan struct{}) <-chan error {
	errs := make(chan error, 2)
	go func() {
		defer close(errs)
		ingress.processRequestQueue(done, errs)
	}()
	return errs
}

func (ingress *ingress) syncFromEpoch(epoch registry.Epoch, pods []registry.Pod) error {
	ingress.podsMu.Lock()
	ingress.podsPrev = ingress.podsCurr
	ingress.podsCurr = map[[32]byte]registry.Pod{}
	for _, pod := range pods {
		ingress.podsCurr[pod.Hash] = pod
	}
	ingress.podsMu.Unlock()
	return nil
}

func (ingress *ingress) processRequestQueue(done <-chan struct{}, errs chan<- error) {
	dispatch.CoForAll(NumBackgroundWorkers, func(i int) {
		for {
			select {
			case <-done:
				return
			case request, ok := <-ingress.queueRequests:
				if !ok {
					return
				}
				switch req := request.(type) {
				case OpenOrderFragmentMappingRequest:
					ingress.processOpenOrderFragmentMappingRequest(req, done, errs)
				default:
					log.Printf("[error] unexpected request type %T", request)
				}
			}
		}
	})
}

func (ingress *ingress) processOpenOrderFragmentMappingRequest(req OpenOrderFragmentMappingRequest, done <-chan struct{}, errs chan<- error) {
	ingress.podsMu.RLock()
	defer ingress.podsMu.RUnlock()

	// Select pods based on the depth
	pods := map[[32]byte]registry.Pod{}
	switch req.orderFragmentEpochDepth {
	case 0:
		pods = ingress.podsCurr
	case 1:
		pods = ingress.podsPrev
	default:
		select {
		case <-done:
		case errs <- fmt.Errorf("[error] (open) order fragment mapping = %v: %v", req.orderID, ErrUnsupportedEpochDepth):
		}
		return
	}

	podDidReceiveFragments := int64(0)

	dispatch.CoForAll(pods, func(hash [32]byte) {
		orderFragments := req.orderFragmentMapping[hash]
		if orderFragments != nil && len(orderFragments) > 0 {
			if err := ingress.sendOrderFragmentsToPod(pods[hash], orderFragments); err != nil {
				select {
				case <-done:
				case errs <- fmt.Errorf("[error] (open) order fragment mapping = %v: %v", req.orderID, err):
				}
				return
			}
			if atomic.LoadInt64(&podDidReceiveFragments) == int64(0) {
				atomic.AddInt64(&podDidReceiveFragments, 1)
			}
		}
	})

	if atomic.LoadInt64(&podDidReceiveFragments) == int64(0) {
		select {
		case <-done:
		case errs <- fmt.Errorf("[error] (open) order fragment mapping = %v: %v", req.orderID, ErrCannotOpenOrderFragments):
		}
		return
	}
}

func (ingress *ingress) sendOrderFragmentsToPod(pod registry.Pod, orderFragments []OrderFragment) error {
	if len(orderFragments) < pod.Threshold() || len(orderFragments) > len(pod.Darknodes) {
		return ErrInvalidNumberOfOrderFragments
	}

	// Map order fragments to their respective Darknodes
	orderFragmentIndexMapping := map[int64]OrderFragment{}
	for _, orderFragment := range orderFragments {
		orderFragmentIndexMapping[orderFragment.Index] = orderFragment
	}

	errs := make(chan error, len(pod.Darknodes))
	go func() {
		defer close(errs)

		logger.Network(logger.LevelInfo, fmt.Sprintf("sending %v order = %v to pod = %v", orderFragments[0].OrderParity, orderFragments[0].OrderID, base64.StdEncoding.EncodeToString(pod.Hash[:8])))

		dispatch.CoForAll(pod.Darknodes, func(i int) {
			orderFragment, ok := orderFragmentIndexMapping[int64(i+1)] // Indices for fragments start at 1
			if !ok {
				errs <- fmt.Errorf("no fragment found at index %v", i)
				return
			}

			darknode := pod.Darknodes[i]

			if len(darknode) == 0 {
				errs <- fmt.Errorf("empty darknode address")
				return
			}

			// Send the order fragment to the Darknode
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
			defer cancel()

			darknodeMultiAddr, err := ingress.swarmer.Query(ctx, darknode)
			if err != nil {
				errs <- fmt.Errorf("cannot send query to %v: %v", darknode, err)
				return
			}

			if err := ingress.orderbookClient.OpenOrder(ctx, darknodeMultiAddr, orderFragment.EncryptedFragment); err != nil {
				log.Printf("cannot send order fragment to %v: %v", darknode, err)
				errs <- fmt.Errorf("cannot send order fragment to %v: %v", darknode, err)
				return
			}
		})
	}()

	// Capture all errors and keep the first error that occurred.
	var errNum int
	var err error
	for errLocal := range errs {
		if errLocal != nil {
			errNum++
			if err == nil {
				err = errLocal
			}
		}
	}

	// Check if at least 2/3 of the nodes in the specified pod have received
	// the order fragments.
	errNumMax := len(orderFragments) - pod.Threshold()
	if len(pod.Darknodes) > 0 && errNum > errNumMax {
		return fmt.Errorf("cannot send order fragments to %v nodes (out of %v nodes) in pod %v: %v", errNum, len(pod.Darknodes), base64.StdEncoding.EncodeToString(pod.Hash[:]), err)
	}
	return nil
}

func (ingress *ingress) verifyOrderFragmentMappings(orderFragmentMappings OrderFragmentMappings) error {
	if len(orderFragmentMappings) == 0 {
		return ErrInvalidOrderFragmentMapping
	}

	ingress.podsMu.RLock()
	defer ingress.podsMu.RUnlock()

	for i := range orderFragmentMappings {
		if err := ingress.verifyOrderFragmentMapping(orderFragmentMappings[i], i); err != nil {
			return err
		}
	}
	return nil
}

func (ingress *ingress) verifyOrderFragmentMapping(orderFragmentMapping OrderFragmentMapping, orderFragmentEpochDepth int) error {
	// Select pods based on the depth
	pods := map[[32]byte]registry.Pod{}
	switch orderFragmentEpochDepth {
	case 0:
		pods = ingress.podsCurr
	case 1:
		pods = ingress.podsPrev
	default:
		return ErrUnsupportedEpochDepth
	}

	if len(orderFragmentMapping) == 0 || len(orderFragmentMapping) > len(pods) {
		logger.Error(fmt.Sprintf("invalid number of pods: got %v, expected %v", len(orderFragmentMapping), len(pods)))
		return ErrInvalidNumberOfPods
	}
	for hash, orderFragments := range orderFragmentMapping {
		pod, ok := pods[hash]
		if !ok {
			return ErrUnknownPod
		}
		if len(orderFragments) > len(pod.Darknodes) || len(orderFragments) < pod.Threshold() {
			return ErrInvalidNumberOfOrderFragments
		}

		// Ensure order fragment Epoch depth matches up with value provided as
		// parameter.
		for _, orderFragment := range orderFragments {
			if orderFragment.EpochDepth != order.FragmentEpochDepth(orderFragmentEpochDepth) {
				return ErrInvalidEpochDepth
			}
		}
	}
	return nil
}

func (ingress *ingress) orderParityFromOrderFragmentMappings(orderFragmentMappings OrderFragmentMappings) order.Parity {
	ingress.podsMu.RLock()
	defer ingress.podsMu.RUnlock()

	for i := range orderFragmentMappings {
		for _, orderFragments := range orderFragmentMappings[i] {
			for _, orderFragment := range orderFragments {
				return orderFragment.OrderParity
			}
		}
	}

	return order.ParityBuy
}
