package ingress

import (
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/republicprotocol/republic-go/order"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type Rewarder interface {
	SelectReferrents(address string) ([]string, error)
	InsertWithdrawalDetails(rewards map[string]*big.Int, hash []byte, address string, token order.Token, amount *big.Int, nonce int64) error
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
		if err := rows.Scan(&sqlReferrent); err != nil {
			return nil, err
		}
		if sqlReferrent.Valid {
			referrents = append(referrents, sqlReferrent.String)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return referrents, nil
}

func (rewarder *rewarder) InsertWithdrawalDetails(rewards map[string]*big.Int, hash []byte, address string, token order.Token, amount *big.Int, nonce int64) error {
	// Ensure user has sufficient rewards
	tokenSymbol := token.String()
	rows, err := rewarder.Query("SELECT amount FROM withdrawals WHERE address=$1 AND token=$2", address, token)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var sqlAmount *big.Int
		if err := rows.Scan(&sqlAmount); err != nil {
			return err
		}
		rewards[tokenSymbol] = new(big.Int).Sub(rewards[tokenSymbol], sqlAmount)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	if amount.Cmp(rewards[tokenSymbol]) == 1 {
		return ErrInsufficientFunds
	}

	// Ensure nonce is valid
	var sqlNonce sql.NullInt64
	if err := rewarder.QueryRow("SELECT nonce FROM withdrawals ORDER BY nonce DESC WHERE address=$1", strings.ToLower(address)).Scan(&nonce); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}
	var previousNonce int64
	if sqlNonce.Valid {
		previousNonce = sqlNonce.Int64
	}
	if previousNonce >= nonce {
		return fmt.Errorf("cannot insert reward details: outdated nonce %v", nonce)
	}

	timestamp := time.Now().Unix()
	_, err = rewarder.Exec("INSERT INTO withdrawals (hash, address, token, amount, timestamp, nonce) VALUES ($1, $2, $3, $4, $5, $6)", hash, strings.ToLower(address), tokenSymbol, amount, timestamp, nonce)
	return err
}
