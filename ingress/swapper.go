package ingress

import (
	"database/sql"
	"encoding/base64"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/republicprotocol/renex-ingress-go/contract"
)

// TABLES
//
// CREATE TABLE partial_swap (
//     order_id        varchar PRIMARY KEY,
//     kyc_addr        varchar,
//     send_to         varchar,
//     receive_from    varchar,
//     time_lock       int,
//     secret_hash     varchar
// );
//
// CREATE TABLE finalized_swap (
//     order_id        	  varchar PRIMARY KEY,
//     send_to               varchar,
//     receive_from          varchar,
//     send_amount           varchar,
//     receive_amount        varchar,
//     secret_hash           varchar,
//     should_initiate_first boolean,
//     time_lock             int
// );

type PartialSwap struct {
	OrderID     string `json:"order_id"`
	KycAddr     string `json:"kyc_addr"`
	SendTo      string `json:"send_to"`
	ReceiveFrom string `json:"receive_from"`
	SecretHash  string `json:"secret_hash"`
	TimeLock    int64  `json:"time_lock"`
}

type FinalizedSwap struct {
	OrderID             string `json:"order_id"`
	SendTo              string `json:"send_to"`
	ReceiveFrom         string `json:"receive_from"`
	SendAmount          string `json:"send_amount"`
	ReceiveAmount       string `json:"receive_amount"`
	SecretHash          string `json:"secret_hash"`
	ShouldInitiateFirst bool   `json:"should_initiate_first"`
	TimeLock            int64  `json:"time_lock"`
}

type Swapper interface {
	InsertPartialSwap(swap PartialSwap) error

	PartialSwap(id string) (PartialSwap, error)

	FinalizedSwap(id string) (FinalizedSwap, bool, error)
}

type swapper struct {
	*sql.DB
	binder contract.Binder
}

func NewSwapper(databaseURL string, binder contract.Binder) (Swapper, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	swapper := &swapper{db, binder}

	return swapper, nil
}

func (swapper *swapper) InsertPartialSwap(swap PartialSwap) error {
	_, err := swapper.Exec("INSERT INTO partial_swap (order_id, kyc_addr, send_to, receive_from ,secret_hash, time_lock) VALUES ($1,$2,$3,$4,$5,$6)",
		swap.OrderID, swap.KycAddr, swap.SendTo, swap.ReceiveFrom, swap.SecretHash, swap.TimeLock)
	return err
}

func (swapper *swapper) PartialSwap(id string) (PartialSwap, error) {
	var swap PartialSwap
	swap.OrderID = id
	err := swapper.QueryRow("SELECT kyc_addr, send_to, receive_from, secret_hash, time_lock FROM partial_swap WHERE order_id = $1", id).
		Scan(&swap.KycAddr, &swap.SendTo, &swap.ReceiveFrom, &swap.SecretHash, &swap.TimeLock)
	return swap, err
}

func (swapper *swapper) FinalizedSwap(id string) (FinalizedSwap, bool, error) {
	orderID, err := orderIdStringToBytes(id)
	if err != nil {
		return FinalizedSwap{}, false, err
	}

	// Check if the order has been canceled
	status, err := swapper.binder.OrderState(orderID)
	if err != nil {
		return FinalizedSwap{}, false, err
	}
	if status == 3 {
		return FinalizedSwap{}, true, nil
	}

	// Get settlement details
	details, err := swapper.binder.GetMatchDetails(orderID)
	if err != nil {
		return FinalizedSwap{}, false, fmt.Errorf("cannot get match details for order=%v, err=%v", id, err)
	}
	if !details.Settled {
		return FinalizedSwap{}, false, fmt.Errorf("order=%v has not been settled", id)
	}

	// Construct the missing fields of the swap.
	var swap FinalizedSwap
	swap.OrderID = id
	pSwap, err := swapper.PartialSwap(swap.OrderID)
	if err != nil {
		return FinalizedSwap{}, false, fmt.Errorf("cannot get partial swap for order=%v, err=%v", swap.OrderID, err)

	}
	matchedID := base64.StdEncoding.EncodeToString(details.MatchedID[:])
	matchedPartialSwap, err := swapper.PartialSwap(matchedID)
	if err != nil {
		return FinalizedSwap{}, false, fmt.Errorf("cannot get matched partial swap for order=%v, err=%v", matchedID, err)
	}
	swap.SendTo = matchedPartialSwap.ReceiveFrom
	swap.ReceiveFrom = matchedPartialSwap.SendTo

	priorityAmount := details.PriorityVolume.String()
	secondaryAmount := details.SecondaryVolume.String()
	if details.OrderIsBuy {
		swap.SendAmount = priorityAmount
		swap.ReceiveAmount = secondaryAmount
		swap.ShouldInitiateFirst = false
		swap.SecretHash = matchedPartialSwap.SecretHash
		swap.TimeLock = matchedPartialSwap.TimeLock
	} else {
		swap.SendAmount = secondaryAmount
		swap.ReceiveAmount = priorityAmount
		swap.ShouldInitiateFirst = true
		swap.SecretHash = pSwap.SecretHash
		swap.TimeLock = pSwap.TimeLock
	}

	return swap, false, nil
}

func orderIdStringToBytes(id string) ([32]byte, error) {
	orderIDBytes, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return [32]byte{}, err
	}
	var orderID [32]byte
	copy(orderID[:], orderIDBytes[:])

	return orderID, nil
}
