package ingress_test

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/republicprotocol/renex-ingress-go/ingress"

	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/identity"
	"github.com/republicprotocol/republic-go/order"
	"github.com/republicprotocol/republic-go/registry"
)

var _ = Describe("Ingress", func() {

	var rsaKey crypto.RsaKey
	var ecdsaKey crypto.EcdsaKey
	var contract ContractBinder
	var renExContract RenExContractBinder
	var ingress Ingress
	var done chan struct{}
	var errChSync <-chan error
	var errChProcess <-chan error

	BeforeEach(func() {
		var err error
		done = make(chan struct{})

		rsaKey, err = crypto.RandomRsaKey()
		Expect(err).ShouldNot(HaveOccurred())

		ecdsaKey, err = crypto.RandomEcdsaKey()
		Expect(err).ShouldNot(HaveOccurred())

		contract = newIngressBinder()

		renExContract = newRenExBinder()

		swarmer := mockSwarmer{}
		orderbookClient := mockOrderbookClient{}

		ingress = NewIngress(ecdsaKey, contract, renExContract, &swarmer, &orderbookClient, time.Millisecond, &mockSwapper{}, &mockLoginer{})
		errChSync = ingress.Sync(done)
		errChProcess = ingress.ProcessRequests(done)

		// Consume errors in the background to allow progress when an event occurs
		go captureErrorsFromErrorChannel(errChSync)
		go captureErrorsFromErrorChannel(errChProcess)

		time.Sleep(100 * time.Millisecond)
	})

	AfterEach(func() {
		close(done)

		// Wait for all errors to close
		captureErrorsFromErrorChannel(errChSync)
		captureErrorsFromErrorChannel(errChProcess)

		time.Sleep(time.Second)
	})

	Context("when approving withdrawals", func() {
		It("should approve valid withdrawals with a valid signature", func() {
			trader := [20]byte{}
			_, err := rand.Read(trader[:])
			Expect(err).ShouldNot(HaveOccurred())

			tokenID := uint32(0)
			// TODO: Retrieve nonce from renExContract (without incrementing it)
			traderNonce := big.NewInt(0)

			signature, err := ingress.ApproveWithdrawal(trader, tokenID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(signature).ShouldNot(BeNil())

			message, err := WithdrawalMessage(trader, tokenID, traderNonce)
			Expect(err).ShouldNot(HaveOccurred())

			signatureData := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))), message)

			broker, err := crypto.RecoverAddress(signatureData, signature[:])
			Expect(err).ShouldNot(HaveOccurred())

			Expect(broker).Should(Equal(ecdsaKey.Address()))
		})
	})

	Context("when opening orders", func() {

		It("should open orders with a sufficient number of order fragments", func() {
			ord, err := createOrder()
			Expect(err).ShouldNot(HaveOccurred())
			fragments, err := ord.Split(6, 4)
			Expect(err).ShouldNot(HaveOccurred())

			pods, err := contract.Pods()
			Expect(err).ShouldNot(HaveOccurred())

			orderFragmentMappingIn := OrderFragmentMapping{}
			orderFragmentMappingIn[pods[0].Hash] = []OrderFragment{}
			for i, fragment := range fragments {
				orderFragment := OrderFragment{
					Index: int64(i),
				}
				orderFragment.EncryptedFragment, err = fragment.Encrypt(rsaKey.PublicKey)
				Expect(err).ShouldNot(HaveOccurred())
				orderFragmentMappingIn[pods[0].Hash] = append(orderFragmentMappingIn[pods[0].Hash], orderFragment)
			}
			orderFragmentMappingsIn := OrderFragmentMappings{orderFragmentMappingIn}

			trader := [20]byte{}
			_, err = rand.Read(trader[:])
			Expect(err).ShouldNot(HaveOccurred())

			signature, err := ingress.OpenOrder(trader, ord.ID, orderFragmentMappingsIn)
			Expect(signature).ShouldNot(BeNil())
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should not open orders with an insufficient number of order fragments", func() {
			ord, err := createOrder()
			Expect(err).ShouldNot(HaveOccurred())
			fragments, err := ord.Split(int64(1), int64(4/3))
			Expect(err).ShouldNot(HaveOccurred())

			orderFragmentMappingIn := OrderFragmentMapping{}
			pods, err := contract.Pods()
			Expect(err).ShouldNot(HaveOccurred())
			orderFragmentMappingIn[pods[0].Hash] = []OrderFragment{}
			for i, fragment := range fragments {
				orderFragment := OrderFragment{
					Index: int64(i),
				}
				orderFragment.EncryptedFragment, err = fragment.Encrypt(rsaKey.PublicKey)
				Expect(err).ShouldNot(HaveOccurred())
				orderFragmentMappingIn[pods[0].Hash] = append(orderFragmentMappingIn[pods[0].Hash], orderFragment)
			}

			orderFragmentMappingsIn := OrderFragmentMappings{}
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)

			trader := [20]byte{}
			_, err = rand.Read(trader[:])
			Expect(err).ShouldNot(HaveOccurred())

			signature, err := ingress.OpenOrder(trader, ord.ID, orderFragmentMappingsIn)
			Expect(signature).ShouldNot(BeNil())
			Expect(err).Should(HaveOccurred())
		})

		It("should not open orders with malformed order fragments", func() {
			ord, err := createOrder()
			Expect(err).ShouldNot(HaveOccurred())

			orderFragmentMappingIn := OrderFragmentMapping{}
			pods, err := contract.Pods()
			Expect(err).ShouldNot(HaveOccurred())
			orderFragmentMappingIn[pods[0].Hash] = make([]OrderFragment, 3)

			orderFragmentMappingsIn := OrderFragmentMappings{}
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)

			trader := [20]byte{}
			_, err = rand.Read(trader[:])
			Expect(err).ShouldNot(HaveOccurred())

			signature, err := ingress.OpenOrder(trader, ord.ID, orderFragmentMappingsIn)
			Expect(signature).ShouldNot(BeNil())
			Expect(err).Should(HaveOccurred())
		})

		It("should not open orders with empty orderFragmentMappings", func() {
			ord, err := createOrder()
			Expect(err).ShouldNot(HaveOccurred())

			orderFragmentMappingIn := OrderFragmentMappings{}

			trader := [20]byte{}
			_, err = rand.Read(trader[:])
			Expect(err).ShouldNot(HaveOccurred())

			signature, err := ingress.OpenOrder(trader, ord.ID, orderFragmentMappingIn)
			Expect(signature).ShouldNot(BeNil())
			Expect(err).Should(HaveOccurred())
		})

		It("should not open orders with unknown pod hashes", func() {
			ord, err := createOrder()
			Expect(err).ShouldNot(HaveOccurred())

			orderFragmentMappingIn := OrderFragmentMapping{}
			orderFragmentMappingIn[[32]byte{byte(1)}] = make([]OrderFragment, 3)

			orderFragmentMappingsIn := OrderFragmentMappings{}
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)

			trader := [20]byte{}
			_, err = rand.Read(trader[:])
			Expect(err).ShouldNot(HaveOccurred())

			signature, err := ingress.OpenOrder(trader, ord.ID, orderFragmentMappingsIn)
			Expect(signature).ShouldNot(BeNil())
			Expect(err).Should(HaveOccurred())
		})

		It("should not open orders with an invalid number of pods", func() {
			ord, err := createOrder()
			Expect(err).ShouldNot(HaveOccurred())
			fragments, err := ord.Split(6, 4)
			Expect(err).ShouldNot(HaveOccurred())

			pods, err := contract.Pods()
			Expect(err).ShouldNot(HaveOccurred())

			orderFragmentMappingIn := OrderFragmentMapping{}
			orderFragmentMappingIn[pods[0].Hash] = []OrderFragment{}
			for i, fragment := range fragments {
				orderFragment := OrderFragment{
					Index: int64(i),
				}
				orderFragment.EncryptedFragment, err = fragment.Encrypt(rsaKey.PublicKey)
				Expect(err).ShouldNot(HaveOccurred())
				orderFragmentMappingIn[pods[0].Hash] = append(orderFragmentMappingIn[pods[0].Hash], orderFragment)
			}

			orderFragmentMappingsIn := OrderFragmentMappings{}
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)

			trader := [20]byte{}
			_, err = rand.Read(trader[:])
			Expect(err).ShouldNot(HaveOccurred())

			signature, err := ingress.OpenOrder(trader, ord.ID, orderFragmentMappingsIn)
			Expect(signature).ShouldNot(BeNil())
			Expect(err).Should(Equal(ErrInvalidNumberOfPods))
		})

		It("should not open orders with an invalid epoch depth", func() {
			ord, err := createOrder()
			Expect(err).ShouldNot(HaveOccurred())
			fragments, err := ord.Split(6, 4)
			Expect(err).ShouldNot(HaveOccurred())

			pods, err := contract.Pods()
			Expect(err).ShouldNot(HaveOccurred())

			orderFragmentMappingIn := OrderFragmentMapping{}
			orderFragmentMappingIn[pods[0].Hash] = []OrderFragment{}
			for i, fragment := range fragments {
				orderFragment := OrderFragment{
					Index: int64(i),
				}
				orderFragment.EncryptedFragment, err = fragment.Encrypt(rsaKey.PublicKey)
				orderFragment.EncryptedFragment.EpochDepth = 2
				Expect(err).ShouldNot(HaveOccurred())
				orderFragmentMappingIn[pods[0].Hash] = append(orderFragmentMappingIn[pods[0].Hash], orderFragment)
			}

			orderFragmentMappingsIn := OrderFragmentMappings{}
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)

			trader := [20]byte{}
			_, err = rand.Read(trader[:])
			Expect(err).ShouldNot(HaveOccurred())

			signature, err := ingress.OpenOrder(trader, ord.ID, orderFragmentMappingsIn)
			Expect(signature).ShouldNot(BeNil())
			Expect(err).Should(Equal(ErrInvalidEpochDepth))
		})
	})
})

// ErrOpenOpenedOrder is returned when trying to open an opened order.
var ErrOpenOpenedOrder = errors.New("cannot open order that is already open")

type renExBinder struct {
	traderNonces map[common.Address]*big.Int
}

func newRenExBinder() *renExBinder {
	return &renExBinder{
		traderNonces: map[common.Address]*big.Int{},
	}
}

func (binder *renExBinder) GetTraderWithdrawalNonce(trader common.Address) (*big.Int, error) {
	nonce := binder.traderNonces[trader]
	if nonce == nil {
		nonce = big.NewInt(0)
	}
	binder.traderNonces[trader] = big.NewInt(0).Add(big.NewInt(1), nonce)
	return nonce, nil
}

func (binder *renExBinder) BalanceOf(common.Address) (*big.Int, error) {
	return big.NewInt(1), nil
}

func (binder *renExBinder) GetOrderTrader(orderID [32]byte) (common.Address, error) {
	return common.Address{}, nil
}

// ingressBinder is a mock implementation of ingress.ContractBinder.
type ingressBinder struct {
	buyOrdersMu *sync.Mutex
	buyOrders   []order.ID

	sellOrdersMu *sync.Mutex
	sellOrders   []order.ID

	ordersMu    *sync.Mutex
	orders      map[order.ID]int
	orderStatus map[order.ID]order.Status

	numberOfDarknodes int
	pods              []registry.Pod
	previousPods      []registry.Pod
}

// newIngressBinder returns a mock ingressBinder.
func newIngressBinder() *ingressBinder {
	pod := registry.Pod{
		Hash:      [32]byte{},
		Darknodes: []identity.Address{},
	}
	rand.Read(pod.Hash[:])
	for i := 0; i < 6; i++ {
		ecdsaKey, err := crypto.RandomEcdsaKey()
		if err != nil {
			panic(fmt.Sprintf("cannot create mock darkpool %v", err))
		}
		pod.Darknodes = append(pod.Darknodes, identity.Address(ecdsaKey.Address()))
	}

	return &ingressBinder{
		buyOrdersMu: new(sync.Mutex),
		buyOrders:   []order.ID{},

		sellOrdersMu: new(sync.Mutex),
		sellOrders:   []order.ID{},

		ordersMu:    new(sync.Mutex),
		orders:      map[order.ID]int{},
		orderStatus: map[order.ID]order.Status{},

		numberOfDarknodes: 6,
		pods:              []registry.Pod{pod},
	}
}

func (binder *ingressBinder) Darknodes() (identity.Addresses, error) {
	darknodes := identity.Addresses{}
	for _, pod := range binder.pods {
		darknodes = append(darknodes, pod.Darknodes...)
	}
	return darknodes, nil
}

func (binder *ingressBinder) NextEpoch() (registry.Epoch, error) {
	return binder.Epoch()
}

func (binder *ingressBinder) PreviousEpoch() (registry.Epoch, error) {
	darknodes, err := binder.Darknodes()
	if err != nil {
		return registry.Epoch{}, err
	}
	return registry.Epoch{
		Hash:      [32]byte{1},
		Pods:      binder.pods,
		Darknodes: darknodes,
	}, nil
}

func (binder *ingressBinder) Epoch() (registry.Epoch, error) {
	darknodes, err := binder.Darknodes()
	if err != nil {
		return registry.Epoch{}, err
	}
	return registry.Epoch{
		Hash:          [32]byte{2},
		Pods:          binder.pods,
		Darknodes:     darknodes,
		BlockNumber:   big.NewInt(0),
		BlockInterval: big.NewInt(1),
	}, nil
}

func (binder *ingressBinder) MinimumEpochInterval() (*big.Int, error) {
	return big.NewInt(1), nil
}

func (binder *ingressBinder) Pods() ([]registry.Pod, error) {
	return binder.pods, nil
}

func (binder *ingressBinder) PreviousPods() ([]registry.Pod, error) {
	return binder.previousPods, nil
}

func (binder *ingressBinder) setOrderStatus(orderID order.ID, status order.Status) error {
	binder.ordersMu.Lock()
	defer binder.ordersMu.Unlock()

	switch status {
	case order.Open:
		binder.orderStatus[orderID] = order.Open
	case order.Confirmed:
		if binder.orderStatus[orderID] != order.Open {
			return errors.New("order not open")
		}
		binder.orderStatus[orderID] = order.Confirmed
	case order.Canceled:
		if binder.orderStatus[orderID] != order.Open {
			return errors.New("order not open")
		}
		binder.orderStatus[orderID] = order.Canceled
	}
	return nil
}

func createOrder() (order.Order, error) {
	parity := order.ParityBuy
	nonce := uint64(mathRand.Intn(1000000000))
	return order.NewOrder(parity, order.TypeLimit, time.Now().Add(time.Hour), order.SettlementRenEx, order.TokensETHREN, 1e12, 1e12, 1e12, nonce), nil
}

type mockSwarmer struct {
}

func (swarmer *mockSwarmer) Ping(ctx context.Context) error {
	return nil
}

func (swarmer *mockSwarmer) Pong(ctx context.Context, to identity.MultiAddress) error {
	return nil
}

func (swarmer *mockSwarmer) BroadcastMultiAddress(ctx context.Context, multiAddress identity.MultiAddress) error {
	return nil
}

func (swarmer *mockSwarmer) Query(ctx context.Context, query identity.Address) (identity.MultiAddress, error) {
	return identity.MultiAddress{}, nil
}

func (swarmer *mockSwarmer) MultiAddress() identity.MultiAddress {
	return identity.MultiAddress{}
}

func (swarmer *mockSwarmer) Peers() (identity.MultiAddresses, error) {
	return identity.MultiAddresses{}, nil
}

type mockOrderbookClient struct {
}

func (client *mockOrderbookClient) OpenOrder(ctx context.Context, to identity.MultiAddress, orderFragment order.EncryptedFragment) error {
	return nil
}

func captureErrorsFromErrorChannel(errs <-chan error) {
	for range errs {
	}
}

type mockSwapper struct {
}

func (swapper *mockSwapper) SelectAuthorizedAddress(kycAddress string) (string, error) {
	return "", nil
}
func (swapper *mockSwapper) InsertAuthorizedAddress(kycAddress string, atomAddress string) error {
	return nil
}
func (swapper *mockSwapper) SelectAddress(orderID string) (string, error) {
	return "", nil
}
func (swapper *mockSwapper) InsertAddress(orderID string, address string) error {
	return nil
}
func (swapper *mockSwapper) SelectSwapDetails(orderID string) (string, error) {
	return "", nil
}
func (swapper *mockSwapper) InsertSwapDetails(orderID string, swapDetails string) error {
	return nil
}

type mockLoginer struct {
}

func (Loginer *mockLoginer) SelectLogin(address string) (int64, string, error) {
	return 0, "", nil
}

func (Loginer *mockLoginer) InsertLogin(address, referrer string) error {
	return nil
}

func (Loginer *mockLoginer) UpdateLogin(address string, kyberUID int64, kycType int) error {
	return nil
}
