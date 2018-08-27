package httpadapter_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"

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

func (adapter *weakAdapter) ApproveWithdrawal(trader string, tokenID uint32) ([65]byte, error) {
	atomic.AddInt64(&adapter.numWithdrawn, 1)
	return WEAK_SIGNATURE, nil
}

type errAdapter struct {
}

func (adapter *errAdapter) OpenOrder(trader string, orderFragmentMapping OrderFragmentMappings) ([65]byte, error) {
	return [65]byte{}, errors.New("cannot open order")
}

func (adapter *errAdapter) ApproveWithdrawal(trader string, tokenID uint32) ([65]byte, error) {
	return [65]byte{}, errors.New("cannot approve withdrawal")
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
			server := NewIngressServer(&adapter)
			server.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusCreated))
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
			server := NewIngressServer(&adapter)
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
			server := NewIngressServer(&adapter)
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
			server := NewIngressServer(&adapter)
			server.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusCreated))
			Expect(UnmarshalSignature(w.Body.String())).To(Equal(WEAK_SIGNATURE))
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
			server := NewIngressServer(&adapter)
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
			server := NewIngressServer(&adapter)
			server.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
