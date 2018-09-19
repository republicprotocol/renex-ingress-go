package ingress

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type KYCer interface {
	SelectTrader(address string) (string, error)
	InsertTrader(address string) error
}

type kycer struct {
	*sql.DB
}

func NewKYCer(databaseURL string) (KYCer, error) {
	db, err := sql.Open("postgres", databaseURL)
	return &kycer{
		db,
	}, err
}

func (kycer *kycer) SelectTrader(address string) (string, error) {
	var trader string
	if err := kycer.QueryRow("SELECT time FROM kyber_traders WHERE address = $1", address).Scan(&trader); err != nil {
		return trader, err
	}
	if trader == "" {
		return trader, fmt.Errorf("requested address not found")
	}
	return trader, nil
}

func (kycer *kycer) InsertTrader(address string) error {
	_, err := kycer.Exec("INSERT INTO kyber_traders (address, created_at, updated_at) VALUES ($1, $2, $3) ON DUPLICATE KEY UPDATE updated_at=$3", address, time.Now(), time.Now())
	return err
}
