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
	SelectLogin(address string) (int64, string, error)
	InsertLogin(address, referrer string) error
	UpdateLogin(address string, kyberUID int64, kycType int) error
	Authorize(authorizer, authorizedAddr string) error
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

func (loginer *loginer) SelectLogin(address string) (int64, string, error) {
	var sqlKyberUID sql.NullInt64
	var sqlTimestamp sql.NullString
	if err := loginer.QueryRow("SELECT kyc_kyber, last_verified_at FROM traders WHERE address=$1", strings.ToLower(address)).Scan(&sqlKyberUID, &sqlTimestamp); err != nil {
		return 0, "", err
	}
	var kyberUID int64
	var timestamp string
	if sqlKyberUID.Valid {
		kyberUID = sqlKyberUID.Int64
	}
	if sqlTimestamp.Valid {
		timestamp = sqlTimestamp.String
	}
	return kyberUID, timestamp, nil
}

func (loginer *loginer) InsertLogin(address, referrer string) error {
	timestamp := time.Now().Unix()
	_, err := loginer.Exec("INSERT INTO traders (address, referrer, referral_code, created_at) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING", strings.ToLower(address), strings.ToLower(referrer), uuid.NewV4().String(), timestamp)
	return err
}

func (loginer *loginer) UpdateLogin(address string, kyberUID int64, kycType int) error {
	timestamp := time.Now().Unix()
	switch kycType {
	case KYCWyre:
		_, err := loginer.Exec("UPDATE traders SET kyc_wyre=$1, last_verified_at=$2 WHERE address=$1", strings.ToLower(address), timestamp)
		return err
	case KYCKyber:
		_, err := loginer.Exec("UPDATE traders SET kyc_kyber=$2, last_verified_at=$3 WHERE address=$1", strings.ToLower(address), kyberUID, timestamp)
		if err != nil {
			return err
		}

		// Use original referral code.
		var referrer sql.NullString
		if err := loginer.QueryRow("SELECT referrer FROM traders WHERE kyc_kyber=$1 ORDER BY created_at LIMIT 1", kyberUID).Scan(&referrer); err != nil {
			return err
		}
		if referrer.Valid {
			_, err = loginer.Exec("UPDATE traders SET referrer=$2 WHERE kyc_kyber=$1", kyberUID, referrer.String)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (loginer *loginer) Authorize(authorizer, authorizedAddr string) error {
	timestamp := time.Now().Unix()
	_, err := loginer.Exec("INSERT INTO traders (address, kyc_wyre, kyc_kyber, authorizer, created_at, last_verified_at) select $1,kyc_wyre,kyc_kyber,$2,$3,last_verified_at FROM traders where address=$2 ON CONFLICT DO NOTHING", strings.ToLower(authorizedAddr), strings.ToLower(authorizer), timestamp)
	return err
}
