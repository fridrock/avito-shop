package storage

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CoinStorage interface {
	SendCoin(int, uuid.UUID, uuid.UUID) error
}

type CoinStorageImpl struct {
	db *sqlx.DB
}

// TODO refactor
func (sc *CoinStorageImpl) SendCoin(amount int, curUserId uuid.UUID, toUserId uuid.UUID) error {
	tx, err := sc.db.Begin()
	if err != nil {
		return err
	}
	//removing coins from current user
	q := `UPDATE users SET coins = coins - $1 WHERE id = $2`
	_, err = tx.Exec(q, amount, curUserId.String())
	if err != nil {
		return err
	}
	//append coins
	q = `UPDATE users SET coins = coins + $1 WHERE id = $2`
	_, err = tx.Exec(q, amount, toUserId.String())
	if err != nil {
		return err
	}

	// creating transaction log
	q = `INSERT INTO coin_transactions(transaction_id, from_id, to_id, amount_of_coins) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(q,
		uuid.New().String(),
		curUserId,
		toUserId,
		amount)
	//rollback if error is not nil
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return errors.Join(err, rollbackErr)
		}
	}
	return tx.Commit()
}

func NewCoinStorage(db *sqlx.DB) CoinStorage {
	return &CoinStorageImpl{
		db: db,
	}
}
