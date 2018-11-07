package ingress

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var (
	ErrAddressNotFound = errors.New("requested address not found")
)

type Loginer interface {
	SelectLogin(address string) (string, error)
	InsertLogin(address, referral string) error
}

type loginer struct {
	*sql.DB
}

func NewLoginer(databaseURL string) (Loginer, error) {
	db, err := sql.Open("postgres", databaseURL)
	return &loginer{
		db,
	}, err
}

func (loginer *loginer) SelectLogin(address string) (string, error) {
	var timestamp string
	if err := loginer.QueryRow("SELECT created_at FROM traders WHERE address = $1", strings.ToLower(address)).Scan(&timestamp); err != nil {
		return timestamp, err
	}
	if timestamp == "" {
		return timestamp, ErrAddressNotFound
	}
	return timestamp, nil
}

func (loginer *loginer) InsertLogin(address, referral string) error {
	timestamp := time.Now().Unix()
	_, err := loginer.Exec("INSERT INTO traders (address, referral, created_at) VALUES ($1, $2, $3)", strings.ToLower(address), strings.ToLower(referral), timestamp)
	return err
}
