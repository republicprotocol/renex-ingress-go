package httpadapter_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"

	"github.com/republicprotocol/renex-ingress-go/ingress"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/republicprotocol/renex-ingress-go/httpadapter"
)

type weakAdapter struct {
	numOpened    int64
	numWithdrawn int64
}

var WEAK_SIGNATURE = [65]byte{'W', 'E', 'A', 'K'}

func (adapter *weakAdapter) OpenOrder(trader string, orderFragmentMapping OrderFragmentMappings) ([65]byte, error) {
	atomic.AddInt64(&adapter.numOpened, 1)
	return WEAK_SIGNATURE, nil
}

func (adapter *weakAdapter) WyreVerified(trader string) (bool, error) {
	return true, nil
}

func (adapter *weakAdapter) ApproveWithdrawal(trader string, tokenID uint32) ([65]byte, error) {
	atomic.AddInt64(&adapter.numWithdrawn, 1)
	return WEAK_SIGNATURE, nil
}

func (adapter *weakAdapter) GetLogin(string) (int64, string, error) {
	return 0, "", nil
}

func (adapter *weakAdapter) PostLogin(string, string) error {
	return nil
}

func (adapter *weakAdapter) PostVerification(string, int64, int) error {
	return nil
}

func (adapter *weakAdapter) Authorize(authorizer, authorizedAddr string) error {
	return nil
}

func (adapter *weakAdapter) InsertPartialSwap(swap ingress.PartialSwap) error {
	return nil
}

func (adapter *weakAdapter) PartialSwap(id string) (ingress.PartialSwap, error) {
	return ingress.PartialSwap{}, nil
}

func (adapter *weakAdapter) FinalizedSwap(id string) (ingress.FinalizedSwap, bool, error) {
	return ingress.FinalizedSwap{}, false, nil
}

type errAdapter struct {
}

func (adapter *errAdapter) OpenOrder(trader string, orderFragmentMapping OrderFragmentMappings) ([65]byte, error) {
	return [65]byte{}, errors.New("cannot open order")
}

func (adapter *errAdapter) WyreVerified(trader string) (bool, error) {
	return false, errors.New("trader not verified")
}

func (adapter *errAdapter) ApproveWithdrawal(trader string, tokenID uint32) ([65]byte, error) {
	return [65]byte{}, errors.New("cannot approve withdrawal")
}

func (adapter *errAdapter) GetLogin(string) (int64, string, error) {
	return 0, "", errors.New("cannot get login")
}

func (adapter *errAdapter) PostLogin(string, string) error {
	return errors.New("cannot post login")
}

func (adapter *errAdapter) PostVerification(string, int64, int) error {
	return errors.New("cannot post verification")
}

func (adapter *errAdapter) Authorize(authorizer, authorizedAddr string) error {
	return nil
}

func (adapter *errAdapter) InsertPartialSwap(swap ingress.PartialSwap) error {
	return nil
}

func (adapter *errAdapter) PartialSwap(id string) (ingress.PartialSwap, error) {
	return ingress.PartialSwap{}, nil
}

func (adapter *errAdapter) FinalizedSwap(id string) (ingress.FinalizedSwap, bool, error) {
	return ingress.FinalizedSwap{}, false, nil
}

var _ = Describe("HTTP handlers", func() {

	Context("when opening orders", func() {

		It("should return status 201 for a valid request", func() {

			mockOrder := new(OpenOrderRequest)
			data, err := json.Marshal(mockOrder)
			Expect(err).ShouldNot(HaveOccurred())

			body := bytes.NewBuffer(data)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://localhost/orders", body)

			adapter := weakAdapter{}
			server := NewIngressServer(&adapter, []string{}, "", "")
			server.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusCreated))

			var response OpenOrderResponse
			err = json.Unmarshal(w.Body.Bytes(), &response)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(UnmarshalSignature(response.Signature)).To(Equal(WEAK_SIGNATURE))

			Expect(atomic.LoadInt64(&adapter.numOpened)).To(Equal(int64(1)))
		})

		It("should return status 400 for an invalid request", func() {

			mockOrder := ""
			data, err := json.Marshal(mockOrder)
			Expect(err).ShouldNot(HaveOccurred())

			body := bytes.NewBuffer(data)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://localhost/orders", body)

			adapter := weakAdapter{}
			server := NewIngressServer(&adapter, []string{}, "", "")
			server.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(atomic.LoadInt64(&adapter.numOpened)).To(Equal(int64(0)))
		})

		It("should return status 500 for ingress adapter errors", func() {

			mockOrder := new(OpenOrderRequest)
			data, err := json.Marshal(mockOrder)
			Expect(err).ShouldNot(HaveOccurred())

			body := bytes.NewBuffer(data)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://localhost/orders", body)

			adapter := errAdapter{}
			server := NewIngressServer(&adapter, []string{}, "", "")
			server.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("when approving withdrawals", func() {

		It("should return status 201 for a valid request", func() {

			mockOrder := new(ApproveWithdrawalRequest)
			data, err := json.Marshal(mockOrder)
			Expect(err).ShouldNot(HaveOccurred())

			body := bytes.NewBuffer(data)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://localhost/withdrawals", body)

			adapter := weakAdapter{}
			server := NewIngressServer(&adapter, []string{}, "", "")
			server.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusCreated))

			var response ApproveWithdrawalResponse
			err = json.Unmarshal(w.Body.Bytes(), &response)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(UnmarshalSignature(response.Signature)).To(Equal(WEAK_SIGNATURE))

			Expect(atomic.LoadInt64(&adapter.numWithdrawn)).To(Equal(int64(1)))
		})

		It("should return status 400 for an invalid request", func() {

			mockOrder := ""
			data, err := json.Marshal(mockOrder)
			Expect(err).ShouldNot(HaveOccurred())

			body := bytes.NewBuffer(data)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://localhost/withdrawals", body)

			adapter := weakAdapter{}
			server := NewIngressServer(&adapter, []string{}, "", "")
			server.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(atomic.LoadInt64(&adapter.numWithdrawn)).To(Equal(int64(0)))
		})

		It("should return status 500 for ingress adapter errors", func() {

			mockOrder := new(ApproveWithdrawalRequest)
			data, err := json.Marshal(mockOrder)
			Expect(err).ShouldNot(HaveOccurred())

			body := bytes.NewBuffer(data)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://localhost/withdrawals", body)

			adapter := errAdapter{}
			server := NewIngressServer(&adapter, []string{}, "", "")
			server.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
