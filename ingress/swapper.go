package ingress

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"github.com/republicprotocol/renex-ingress-go/contract/bindings"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	settlement contract.RenExSettlement
	orderbook  bindings.Orderbook
}

func NewSwapper(databaseURL string, settlement contract.RenExSettlement, orderbook bindings.Orderbook) (Swapper, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil ,err
	}
	swapper := &swapper{db, settlement, orderbook}
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
	orderSettled := make(chan *contract.RenExSettlementLogOrderSettled)
	sub, err := swapper.settlement.WatchLogOrderSettled(&bind.WatchOpts{}, orderSettled, orderIDs)
	if err != nil {
		log.Println("cannot subscribe to the contract", err )
		return
	}
	defer sub.Unsubscribe()
	for notification := range orderSettled {
		var buyID, sellID [32]byte
		details, err  := swapper.settlement.GetMatchDetails(&bind.CallOpts{}, notification.OrderID)
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

		buyer, err := swapper.orderbook.OrderTrader(&bind.CallOpts{}, buyID )
		if err != nil {
			log.Printf("cannot get buyer address for order=%v, err=%v", buyID, err)
			continue
		}
		seller, err := swapper.orderbook.OrderTrader(&bind.CallOpts{}, sellID )
		if err != nil {
			log.Printf("cannot get seller address for order=%v, err=%v", sellID, err)
			continue
		}
		buyerAtomicAddr, err  := swapper.SelectAuthorizedAddress(buyer.Hex())
		if err != nil {
			log.Printf("cannot get atomic address for trader=%v, err=%v", buyer.Hex(), err)
			continue
		}
		sellerAtomicAddr, err  := swapper.SelectAuthorizedAddress(seller.Hex())
		if err != nil {
			log.Printf("cannot get atomic address for trader=%v, err=%v", seller.Hex(), err)
			continue
		}

		var pswap PartialSwap
		if details.OrderIsBuy {
			pswap.ReceiveAmount = details.PriorityVolume.String()
			pswap.SendAmount = details.SecondaryVolume.String()
			pswap.SendTo = seller.Hex()
			pswap.ReceiveFrom = sellerAtomicAddr
		} else {
			pswap.ReceiveAmount = details.SecondaryVolume.String()
			pswap.SendAmount = details.PriorityVolume.String()
			pswap.SendTo = buyerAtomicAddr
			pswap.ReceiveFrom = buyer.Hex()
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
	}
}

// // Filter contract's event log
// previousOwners, newOwners := make([]common.Address, 0), make([]common.Address, 0)
// // newOwners = append(newOwners, common.HexToAddress("0xFd974e09363F7F823Ce360d2a2006733AEb3e297"))
// iter, err := orderbook.FilterOwnershipTransferred(&bind.FilterOpts{}, previousOwners, newOwners)
// if err != nil {
// 	log.Fatal(err)
// }
//
// for iter.Next() {
// 	log.Println("previous owner: ", iter.Event.PreviousOwner.Hex(), ", new owner: ", iter.Event.NewOwner.Hex())
// }
