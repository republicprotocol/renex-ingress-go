package ingress

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
)

const (
	KYCNone  int = 0
	KYCWyre  int = 1
	KYCKyber int = 2
)

type Loginer interface {
	SelectLogin(address string) (string, string, error)
	InsertLogin(address, referrer string) error
	UpdateLogin(address, uID string, kycType int) error
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

func (loginer *loginer) SelectLogin(address string) (string, string, error) {
	var kyberUID string
	var timestamp string
	if err := loginer.QueryRow("SELECT kyc_kyber, updated_at FROM traders WHERE address=$1", strings.ToLower(address)).Scan(&kyberUID, &timestamp); err != nil {
		return "", "", err
	}
	return kyberUID, timestamp, nil
}

func (loginer *loginer) InsertLogin(address, referrer string) error {
	timestamp := time.Now().Unix()
	_, err := loginer.Exec("INSERT INTO traders (address, referrer, referral_code, created_at, updated_at) VALUES ($1, $2, $3, $4, $4)", strings.ToLower(address), strings.ToLower(referrer), uuid.NewV4().String(), timestamp)
	return err
}

func (loginer *loginer) UpdateLogin(address, uID string, kycType int) error {
	timestamp := time.Now().Unix()
	switch kycType {
	case KYCWyre:
		_, err := loginer.Exec("UPDATE traders SET kyc_wyre=$2, updated_at=$3 WHERE address=$1", strings.ToLower(address), strings.ToLower(uID), timestamp)
		return err
	case KYCKyber:
		_, err := loginer.Exec("UPDATE traders SET kyc_kyber=$2, updated_at=$3 WHERE address=$1", strings.ToLower(address), strings.ToLower(uID), timestamp)
		if err != nil {
			return err
		}

		// Use original referral code.
		var referrer string
		if err := loginer.QueryRow("SELECT referrer FROM traders WHERE kyc_kyber=$1 ORDER BY created_at LIMIT 1", strings.ToLower(uID)).Scan(&referrer); err != nil {
			return err
		}
		_, err = loginer.Exec("UPDATE traders SET referrer=$2 WHERE kyc_kyber=$1", strings.ToLower(uID), referrer)
		if err != nil {
			return err
		}
	}
	return nil
}
