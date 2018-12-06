package ingress

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"log"
	"strings"

	_ "github.com/lib/pq"
	"github.com/republicprotocol/renex-ingress-go/contract"
)

// TABLES
//
// CREATE TABLE swaps (
//     order_id        varchar NOT NULL,
//     address         varchar,
//     swap_details    varchar
// );

// CREATE TABLE auth_addresses (
//     address      varchar(42) NOT NULL,
//     atom_address varchar
// );

type PartialSwap struct {
	SendTo         string `json:"send_to"`
	ReceiveFrom  string `json:"receive_from"`
	SendAmount string `json:"send_amount"`
	ReceiveAmount string `json:"receive_amount"`
}

type Swapper interface {
	SelectAuthorizedAddress(kycAddress string) (string, error)
	InsertAuthorizedAddress(kycAddress string, atomAddress string) error
	SelectAddress(orderID string) (string, error)
	InsertAddress(orderID string, address string) error
	SelectSwapDetails(orderID string) (string, error)
	InsertSwapDetails(orderID string, swapDetails string) error
}

type swapper struct {
	*sql.DB
	binder contract.Binder
}

func NewSwapper(databaseURL string,  binder contract.Binder) (Swapper, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil ,err
	}
	swapper := &swapper{db, binder}
	go swapper.syncSettlement()
	return swapper, nil
}

func (swapper *swapper) SelectAddress(orderID string) (string, error) {
	var address string
	err := swapper.QueryRow("SELECT address FROM swaps WHERE order_id = $1", orderID).Scan(&address)
	return address, err
}

func (swapper *swapper) InsertAddress(orderID string, address string) error {
	_, err := swapper.Exec("INSERT INTO swaps (order_id, address) VALUES ($1,$2)", orderID, address)
	return err
}

func (swapper *swapper) SelectSwapDetails(orderID string) (string, error) {
	var swapDetails string
	err := swapper.QueryRow("SELECT swap_details FROM swaps WHERE order_id = $1", orderID).Scan(&swapDetails)
	return swapDetails, err
}

func (swapper *swapper) InsertSwapDetails(orderID string, swapDetails string) error {
	_, err := swapper.Exec("INSERT INTO swaps (order_id, swap_details) VALUES ($1,$2)", orderID, swapDetails)
	return err
}

func (swapper *swapper) SelectAuthorizedAddress(kycAddress string) (string, error) {
	var authorizedAddress string
	err := swapper.QueryRow("SELECT atom_address FROM auth_addresses WHERE address = $1", strings.ToLower(kycAddress)).Scan(&authorizedAddress)
	return authorizedAddress, err
}

func (swapper *swapper) InsertAuthorizedAddress(kycAddress, authorizedAddress string) error {
	_, err := swapper.Exec("INSERT INTO auth_addresses (address, atom_address) VALUES ($1,$2) ON CONFLICT (address) DO UPDATE SET atom_address = EXCLUDED.atom_address;", strings.ToLower(kycAddress), strings.ToLower(authorizedAddress))
	return err
}

// syncSettlement listens to the order settlement event from the settlement
// contract. If the settled order is an atomic swap order, we collect the data
// needed for the swap and store them in the database.
func (swapper *swapper) syncSettlement() {
	orderIDs := make([][32]byte, 0)
	orderSettled, err := swapper.binder.WatchLogOrderSettled(orderIDs)
	if err != nil {
		log.Println("cannot subscribe to the contract", err )
		return
	}
	for notification := range orderSettled {
		log.Println("have new notification", hex.EncodeToString(notification.OrderID[:]), )
		var buyID, sellID [32]byte
		details, err  := swapper.binder.GetMatchDetails(notification.OrderID)
		if err != nil {
			log.Printf("cannot get match details for order=%v, err=%v", hex.EncodeToString(notification.OrderID[:]), err)
			continue
		}
		if details.PriorityToken != 0{
			continue
		}

		if details.OrderIsBuy{
			buyID, sellID = notification.OrderID, details.MatchedID
		} else {
			buyID, sellID = details.MatchedID, notification.OrderID
		}

		buyer, err := swapper.binder.OrderTrader(buyID)
		if err != nil {
			log.Printf("cannot get buyer address for order=%v, err=%v", buyID, err)
			continue
		}
		seller, err := swapper.binder.OrderTrader(sellID)
		if err != nil {
			log.Printf("cannot get seller address for order=%v, err=%v", sellID, err)
			continue
		}
		buyerAtomicAddr, err  := swapper.SelectAuthorizedAddress(buyer)
		if err != nil {
			log.Printf("cannot get atomic address for trader=%v, err=%v", buyer, err)
			continue
		}
		sellerAtomicAddr, err  := swapper.SelectAuthorizedAddress(seller)
		if err != nil {
			log.Printf("cannot get atomic address for trader=%v, err=%v", seller, err)
			continue
		}

		var pswap PartialSwap
		if details.OrderIsBuy {
			pswap.ReceiveAmount = details.PriorityVolume.String()
			pswap.SendAmount = details.SecondaryVolume.String()
			pswap.SendTo = seller
			pswap.ReceiveFrom = sellerAtomicAddr
		} else {
			pswap.ReceiveAmount = details.SecondaryVolume.String()
			pswap.SendAmount = details.PriorityVolume.String()
			pswap.SendTo = buyerAtomicAddr
			pswap.ReceiveFrom = buyer
		}

		data, err := json.Marshal(pswap)
		if err != nil{
			log.Printf("cannot marshal the swap detail, %v", err)
			continue
		}
		if err := swapper.InsertSwapDetails(hex.EncodeToString(notification.OrderID[:]), string(data)); err != nil {
			log.Printf("cannot insert swap details for order=%v, err=%v",hex.EncodeToString(notification.OrderID[:]),  err)
			continue
		}
		log.Println("inserted into the db")
	}
}
