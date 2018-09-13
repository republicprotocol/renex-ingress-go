package httpadapter_test

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	mathRand "math/rand"
	"sync/atomic"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/republicprotocol/renex-ingress-go/httpadapter"

	"github.com/republicprotocol/renex-ingress-go/ingress"
	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/order"
)

var _ = Describe("Ingress Adapter", func() {

	Context("when marshaling and unmarshaling order fragment mappings", func() {

		var ord order.Order
		var orderFragmentMappingIn OrderFragmentMapping
		var podHashBytes [32]byte
		var podHash string

		BeforeEach(func() {
			rsaKey, err := crypto.RandomRsaKey()
			Expect(err).ShouldNot(HaveOccurred())
			ord, err = createOrder()
			Expect(err).ShouldNot(HaveOccurred())
			fragments, err := ord.Split(24, 16)
			Expect(err).ShouldNot(HaveOccurred())

			signatureBytes := [65]byte{}
			_, err = rand.Read(signatureBytes[:])
			Expect(err).ShouldNot(HaveOccurred())

			podHashBytes = [32]byte{}
			_, err = rand.Read(podHashBytes[:])
			Expect(err).ShouldNot(HaveOccurred())
			podHash = base64.StdEncoding.EncodeToString(podHashBytes[:])

			orderFragmentMappingIn = OrderFragmentMapping{}
			orderFragmentMappingIn[podHash] = []OrderFragment{}
			for i, fragment := range fragments {
				orderFragment := ingress.OrderFragment{
					Index: int64(i),
				}
				orderFragment.EncryptedFragment, err = fragment.Encrypt(rsaKey.PublicKey)
				Expect(err).ShouldNot(HaveOccurred())
				orderFragmentMappingIn[podHash] = append(
					orderFragmentMappingIn[podHash],
					MarshalOrderFragment(orderFragment))
			}
		})

		It("should return the same data after marshaling and unmarshaling well formed data", func() {
			ordID, orderFragmentMapping, err := UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ordID).Should(Equal(ord.ID))
			Expect(orderFragmentMapping).Should(HaveLen(1))

			for i, fragment := range orderFragmentMapping[podHashBytes] {
				orderFragmentIn := MarshalOrderFragment(fragment)
				Expect(orderFragmentIn).Should(Equal(orderFragmentMappingIn[podHash][i]))
			}
		})

		It("should return an error for malformed order fragment IDs", func() {
			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].ID = orderFragmentMappingIn[podHash][i].ID[1:]
			}
			_, _, err := UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(HaveOccurred())

			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].ID = "this is invalid"
			}
			_, _, err = UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(HaveOccurred())
		})

		It("should return an error for malformed order fragment IDs", func() {
			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].OrderID = orderFragmentMappingIn[podHash][i].OrderID[1:]
			}
			_, _, err := UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(HaveOccurred())

			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].OrderID = "this is invalid"
			}
			_, _, err = UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(HaveOccurred())
		})

		It("should return an error for malformed pod hashes", func() {
			orderFragmentMappingIn[podHash[16:]] = orderFragmentMappingIn[podHash]
			_, _, err := UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(MatchError(ErrInvalidPodHashLength))

			delete(orderFragmentMappingIn, podHash[16:])
			orderFragmentMappingIn["this is invalid"] = orderFragmentMappingIn[podHash]
			_, _, err = UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(HaveOccurred())
		})

		It("should return an error for malformed tokens", func() {
			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].Tokens = "this is invalid"
			}
			_, _, err := UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(HaveOccurred())
		})

		It("should return an error for malformed prices", func() {
			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].Price = []string{}
			}
			_, _, err := UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(MatchError(ErrInvalidEncryptedCoExpShareLength))

			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].Price = []string{"this is invalid", "this is also invalid"}
			}
			_, _, err = UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(HaveOccurred())
		})

		It("should return an error for malformed volumes", func() {
			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].Volume = []string{}
			}
			_, _, err := UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(MatchError(ErrInvalidEncryptedCoExpShareLength))

			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].Volume = []string{"this is invalid", "this is also invalid"}
			}
			_, _, err = UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(HaveOccurred())
		})

		It("should return an error for malformed minimum volumes", func() {
			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].MinimumVolume = []string{}
			}
			_, _, err := UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(MatchError(ErrInvalidEncryptedCoExpShareLength))

			for i := range orderFragmentMappingIn[podHash] {
				orderFragmentMappingIn[podHash][i].MinimumVolume = []string{"this is invalid", "this is also invalid"}
			}
			_, _, err = UnmarshalOrderFragmentMapping(orderFragmentMappingIn)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("when opening orders", func() {

		It("should forward data to the ingress if the signature and mapping are well formed", func() {
			ingress := &mockIngress{&mockSwapper{}, 0, 0}
			ingressAdapter := NewIngressAdapter(ingress)

			traderBytes := [20]byte{}
			_, err := rand.Read(traderBytes[:])
			Expect(err).ShouldNot(HaveOccurred())
			trader := hex.EncodeToString(traderBytes[:])

			orderFragmentMappingIn := OrderFragmentMapping{}
			orderFragmentMappingIn["Td2YBy0MRYPYqqBduRmDsIhTySQUlMhPBM+wnNPWKqq="] = []OrderFragment{}

			orderFragmentMappingsIn := OrderFragmentMappings{}
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)
			_, err = ingressAdapter.OpenOrder(trader, orderFragmentMappingsIn)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(atomic.LoadInt64(&ingress.numOpened)).To(Equal(int64(1)))
		})

		It("should not call ingress.OpenOrder if trader is invalid", func() {
			ingress := &mockIngress{&mockSwapper{}, 0, 0}
			ingressAdapter := NewIngressAdapter(ingress)
			traderBytes := []byte{}
			copy(traderBytes[:], "incorrect trader")
			orderFragmentMappingIn := OrderFragmentMapping{}
			orderFragmentMappingIn["Td2YBy0MRYPYqqBduRmDsIhTySQUlMhPBM+wnNPWKqq="] = []OrderFragment{}

			orderFragmentMappingsIn := OrderFragmentMappings{}
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)

			_, err := ingressAdapter.OpenOrder(string(traderBytes), orderFragmentMappingsIn)
			Expect(err).Should(MatchError(ErrInvalidAddressLength))
			Expect(atomic.LoadInt64(&ingress.numOpened)).To(Equal(int64(0)))
		})

		It("should not call ingress.OpenOrder if pool hash is invalid", func() {
			ingress := &mockIngress{&mockSwapper{}, 0, 0}
			ingressAdapter := NewIngressAdapter(ingress)
			traderBytes := [20]byte{}
			_, err := rand.Read(traderBytes[:])
			Expect(err).ShouldNot(HaveOccurred())
			trader := hex.EncodeToString(traderBytes[:])
			orderFragmentMappingIn := OrderFragmentMapping{}
			orderFragmentMappingIn["some invalid hash"] = []OrderFragment{OrderFragment{OrderID: "thisIsAnOrderID"}}

			orderFragmentMappingsIn := OrderFragmentMappings{}
			orderFragmentMappingsIn = append(orderFragmentMappingsIn, orderFragmentMappingIn)

			_, err = ingressAdapter.OpenOrder(trader, orderFragmentMappingsIn)
			Expect(err).Should(HaveOccurred())
			Expect(atomic.LoadInt64(&ingress.numOpened)).To(Equal(int64(0)))
		})
	})

	Context("when approving withdrawals", func() {

		It("should forward data to the ingress if the signature and mapping are well formed", func() {
			ingress := &mockIngress{&mockSwapper{}, 0, 0}
			ingressAdapter := NewIngressAdapter(ingress)

			traderBytes := [20]byte{}
			_, err := rand.Read(traderBytes[:])
			Expect(err).ShouldNot(HaveOccurred())
			trader := hex.EncodeToString(traderBytes[:])

			_, err = ingressAdapter.ApproveWithdrawal(trader, 0)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(atomic.LoadInt64(&ingress.numWithdrawn)).To(Equal(int64(1)))
		})

		It("should not call ingress.ApproveWithdrawal if trader is invalid", func() {
			ingress := &mockIngress{&mockSwapper{}, 0, 0}
			ingressAdapter := NewIngressAdapter(ingress)
			traderBytes := []byte{}
			copy(traderBytes[:], "incorrect trader")

			_, err := ingressAdapter.ApproveWithdrawal(string(traderBytes), 0)
			Expect(err).Should(MatchError(ErrInvalidAddressLength))
			Expect(atomic.LoadInt64(&ingress.numWithdrawn)).To(Equal(int64(0)))
		})
	})
})

type mockSwapper struct {
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

type mockIngress struct {
	ingress.Swapper
	numOpened    int64
	numWithdrawn int64
}

func (ingress *mockIngress) Sync(done <-chan struct{}) <-chan error {
	return nil
}

func (ingress *mockIngress) OpenOrder(address [20]byte, orderID order.ID, orderFragmentMappings ingress.OrderFragmentMappings) ([65]byte, error) {
	atomic.AddInt64(&ingress.numOpened, 1)
	return [65]byte{}, nil
}

func (ingress *mockIngress) TraderVerified(address [20]byte) (bool, error) {
	return true, nil
}

func (ingress *mockIngress) ApproveWithdrawal(trader [20]byte, tokenID uint32) ([65]byte, error) {
	atomic.AddInt64(&ingress.numWithdrawn, 1)
	return [65]byte{}, nil
}

func (ingress *mockIngress) ProcessRequests(done <-chan struct{}) <-chan error {
	return nil
}

func createOrder() (order.Order, error) {
	parity := order.ParityBuy
	nonce := uint64(mathRand.Intn(1000000000))
	return order.NewOrder(parity, order.TypeLimit, time.Now().Add(time.Hour), order.SettlementRenEx, order.TokensETHREN, 1e12, 1e12, 1e12, nonce), nil
}
