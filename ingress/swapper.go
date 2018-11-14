package ingress

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

// TABLES
//
// CREATE TABLE swaps (
//     order_id         varchar NOT NULL,
//     address         varchar,
//     swap_details varchar
// );

// CREATE TABLE auth_addresses (
//     address           varchar(42) NOT NULL,
//     atom_address varchar
// );

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
}

func NewSwapper(databaseURL string) (Swapper, error) {
	db, err := sql.Open("postgres", databaseURL)
	return &swapper{
		db,
	}, err
}

func (swapper *swapper) SelectAddress(orderID string) (string, error) {
	var address string
	if err := swapper.QueryRow("SELECT address FROM swaps WHERE order_id = $1", orderID).Scan(&address); err != nil {
		return address, err
	}
	if address == "" {
		return address, fmt.Errorf("requested address not found")
	}
	return address, nil
}

func (swapper *swapper) InsertAddress(orderID string, address string) error {
	_, err := swapper.Exec("INSERT INTO swaps (order_id, address) VALUES ($1,$2)", orderID, address)
	return err
}

func (swapper *swapper) SelectSwapDetails(orderID string) (string, error) {
	var swapDetails string
	if err := swapper.QueryRow("SELECT swap_details FROM swaps WHERE order_id = $1", orderID).Scan(&swapDetails); err != nil {
		return "", err
	}
	if swapDetails == "" {
		return swapDetails, fmt.Errorf("requested swap details not found")
	}
	return swapDetails, nil
}

func (swapper *swapper) InsertSwapDetails(orderID string, swapDetails string) error {
	_, err := swapper.Exec("INSERT INTO swaps (order_id, swap_details) VALUES ($1,$2)", orderID, swapDetails)
	return err
}

func (swapper *swapper) SelectAuthorizedAddress(kycAddress string) (string, error) {
	var authorizedAddress string
	if err := swapper.QueryRow("SELECT atom_address FROM auth_addresses WHERE address = $1", kycAddress).Scan(&authorizedAddress); err != nil {
		return "", err
	}
	if authorizedAddress == "" {
		return authorizedAddress, fmt.Errorf("requested authorized address not found")
	}
	return authorizedAddress, nil
}

func (swapper *swapper) InsertAuthorizedAddress(kycAddress, authorizedAddress string) error {
	_, err := swapper.Exec("INSERT INTO auth_addresses (address, atom_address) VALUES ($1,$2) ON CONFLICT (address) DO UPDATE SET atom_address = EXCLUDED.atom_address;", strings.ToLower(kycAddress), strings.ToLower(authorizedAddress))
	return err
}
