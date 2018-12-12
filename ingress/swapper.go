package ingress

import (
	"database/sql"
	"encoding/hex"
	"log"
	"time"

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

	FinalizedSwap(id string) (FinalizedSwap, error)
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
	go swapper.syncSettlement()
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

func (swapper *swapper) FinalizedSwap(id string) (FinalizedSwap, error) {
	var swap FinalizedSwap
	swap.OrderID = id
	err := swapper.QueryRow("SELECT send_to, receive_from, send_amount, receive_amount, secret_hash, should_initiate_first, time_lock FROM finalized_swap WHERE order_id = $1", id).
		Scan(&swap.SendTo, &swap.ReceiveFrom, &swap.SendAmount, &swap.ReceiveAmount, &swap.SecretHash, &swap.ShouldInitiateFirst, &swap.TimeLock)
	return swap, err
}

func (swapper *swapper) insertFinalizedSwap(swap FinalizedSwap) error {
	_, err := swapper.Exec("INSERT INTO partial_swap (order_id, send_to, receive_from, send_amount, receive_amount,secret_hash, should_initiate_first, time_lock) VALUES ($1,$2,$3,$4,$5,$6,$7,&8)",
		swap.OrderID, swap.SendTo, swap.ReceiveFrom, swap.SendAmount, swap.ReceiveAmount, swap.SecretHash, swap.ShouldInitiateFirst, swap.TimeLock)
	return err
}

// syncSettlement listens to the order settlement event from the settlement
// contract. If the settled order is an atomic swap order, we collect the data
// needed for the swap and store them in the database.
func (swapper *swapper) syncSettlement() {
	orderIDs := make([][32]byte, 0)
	orderSettled, err := swapper.binder.WatchLogOrderSettled(orderIDs)
	if err != nil {
		log.Println("cannot subscribe to the contract", err)
		return
	}
	log.Println("start syncing notification from settlement...")
	for notification := range orderSettled {
		log.Println("have new notification", hex.EncodeToString(notification.OrderID[:]))
		details, err := swapper.binder.GetMatchDetails(notification.OrderID)
		if err != nil {
			log.Printf("cannot get match details for order=%v, err=%v", hex.EncodeToString(notification.OrderID[:]), err)
			continue
		}
		if details.PriorityToken != 0 || !details.Settled {
			continue
		}

		var swap FinalizedSwap
		swap.OrderID = hex.EncodeToString(notification.OrderID[:])
		pSwap, err := swapper.PartialSwap(swap.OrderID)
		if err != nil {
			log.Printf("cannot get partial swap for order=%v, err=%v", swap.OrderID, err)
			continue
		}
		matchedID := hex.EncodeToString(details.MatchedID[:])
		matchedPartialSwap, err := swapper.PartialSwap(matchedID)
		if err != nil {
			log.Printf("cannot get matched partial swap for order=%v, err=%v", matchedID, err)
			continue
		}
		timelock := time.Now().Add(48 * time.Hour)
		swap.SendTo = matchedPartialSwap.ReceiveFrom
		swap.ReceiveFrom = matchedPartialSwap.SendTo
		swap.TimeLock = timelock.Unix()

		priorityAmount := details.PriorityVolume.String()
		secondaryAmount := details.SecondaryVolume.String()
		if details.OrderIsBuy {
			swap.SendAmount = secondaryAmount
			swap.ReceiveAmount = priorityAmount
			swap.ShouldInitiateFirst = false
			swap.SecretHash = matchedPartialSwap.SecretHash
		} else {
			swap.SendAmount = priorityAmount
			swap.ReceiveAmount = secondaryAmount
			swap.ShouldInitiateFirst = true
			swap.SecretHash = pSwap.SecretHash
		}

		if err := swapper.insertFinalizedSwap(swap); err != nil {
			log.Printf("cannot get insert finalized swap into database, order=%v, err=%v", swap.OrderID, err)
		}
	}
}
