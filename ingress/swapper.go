package ingress

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Swapper interface {
	InsertAuthorizedAddress(kycAddress string, atomAddress string) error
	GetAuthorizedAddress(kycAddress string) (string, error)
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
	if err := swapper.QueryRow("SELECT address FROM swaps WHERE orderID = $1", orderID).Scan(&address); err != nil {
		return address, err
	}
	if address == "" {
		return address, fmt.Errorf("Requested address not found")
	}
	return address, nil
}

func (swapper *swapper) InsertAddress(orderID string, address string) error {
	_, err := swapper.Exec("INSERT INTO swaps (orderID, address) VALUES ($1,$2)", orderID, address)
	return err
}

func (swapper *swapper) SelectSwapDetails(orderID string) (string, error) {
	var swapDetails string
	if err := swapper.QueryRow("SELECT swapDetails FROM swaps WHERE orderID = $1", orderID).Scan(&swapDetails); err != nil {
		return "", err
	}
	if swapDetails == "" {
		return swapDetails, fmt.Errorf("Requested swap details not found")
	}
	return swapDetails, nil
}

func (swapper *swapper) InsertSwapDetails(orderID string, swapDetails string) error {
	_, err := swapper.Exec("INSERT INTO swaps (orderID, swapDetails) VALUES ($1,$2)", orderID, swapDetails)
	return err
}

func (swapper *swapper) GetAuthorizedAddress(kycAddress string) (string, error) {
	var authorizedAddress string
	if err := swapper.QueryRow("SELECT authorizedAddresses FROM swaps WHERE kycAddress = $1", kycAddress).Scan(&authorizedAddress); err != nil {
		return "", err
	}
	if authorizedAddress == "" {
		return authorizedAddress, fmt.Errorf("Requested authorized address not found")
	}
	return authorizedAddress, nil
}

func (swapper *swapper) InsertAuthorizedAddress(kycAddress, authorizedAddress string) error {
	_, err := swapper.Exec("INSERT INTO authorizedAddresses (kycAddress, authorizedAddress) VALUES ($1,$2)", kycAddress, authorizedAddress)
	return err
}
