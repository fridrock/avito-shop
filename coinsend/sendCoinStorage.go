package coinsend

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/fridrock/avito-shop/api"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var WrongData = fmt.Errorf("wrong data")

type SendCoinStorage interface {
	SendCoin(api.SendCoinRequest, uuid.UUID) error
}

type SendCoinStorageImpl struct {
	db *sqlx.DB
}

func (sc *SendCoinStorageImpl) SendCoin(sendCoinRequest api.SendCoinRequest, curUserId uuid.UUID) error {
	tx, err := sc.db.Begin()
	if err != nil {
		return err
	}
	//removing coins from current user
	q := `UPDATE users SET coins = coins - $1 WHERE id = $2 AND coins - $1 > 0`
	_, err = tx.Exec(q, sendCoinRequest.Amount, curUserId.String())
	if err != nil {
		return WrongData
	}
	//getting user to append coins
	q = `SELECT id FROM users WHERE username = $1`
	var toUserId uuid.UUID
	err = tx.QueryRow(q, sendCoinRequest.ToUser).Scan(&toUserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return WrongData
		}
	}
	//append coins
	q = `UPDATE users SET coins = coins + $1 WHERE id = $2`
	_, err = tx.Exec(q, sendCoinRequest.Amount, toUserId.String())
	if err != nil {
		return fmt.Errorf("unexpected error: %v", err)
	}
	// creating transaction log
	q = `INSERT INTO coin_transactions(transaction_id, from_id, to_id, amount_of_coins) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(q,
		uuid.New().String(),
		curUserId,
		toUserId,
		sendCoinRequest.Amount)
	//rollback if error is not nil
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return errors.Join(err, rollbackErr)
		}
	}
	return tx.Commit()
}

func NewSendCoinStorage(db *sqlx.DB) SendCoinStorage {
	return &SendCoinStorageImpl{
		db: db,
	}
}
