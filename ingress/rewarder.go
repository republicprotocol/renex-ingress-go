package ingress

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
)

type Rewarder interface {
	SelectReferrents(address string) ([]string, error)
}

type rewarder struct {
	*sql.DB
}

func NewRewarder(databaseURL string) (Rewarder, error) {
	db, err := sql.Open("postgres", databaseURL)
	return &rewarder{
		db,
	}, err
}

func (rewarder *rewarder) SelectReferrents(address string) ([]string, error) {
	// Select traders referral code
	var sqlReferrer sql.NullString
	if err := rewarder.QueryRow("SELECT referral_code FROM traders WHERE address=$1", strings.ToLower(address)).Scan(&sqlReferrer); err != nil {
		return nil, err
	}
	var referrer string
	if sqlReferrer.Valid {
		referrer = sqlReferrer.String
	}

	// Find traders that are using this code
	rows, err := rewarder.Query("SELECT address FROM traders WHERE referrer=$1", referrer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var referrents []string
	for rows.Next() {
		var sqlReferrent sql.NullString
		err = rows.Scan(&sqlReferrent)
		if err != nil {
			return nil, err
		}
		if sqlReferrent.Valid {
			referrents = append(referrents, sqlReferrent.String)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return referrents, nil
}
