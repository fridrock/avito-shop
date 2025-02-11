package storage

import (
	"database/sql"
	"errors"

	"github.com/fridrock/avito-shop/api"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const getBoughtsQuery = `SELECT products.product_name as type, COUNT(*) as quantity FROM public.boughts
            INNER JOIN products ON products.id = product_id
            WHERE user_id = $1
            GROUP BY products.product_name;`
const getSentCoinsQuery = `SELECT users.username as toUser, SUM(amount_of_coins) as amount FROM public.coin_transactions
			INNER JOIN users ON users.id = to_id
			WHERE from_id = $1
			GROUP BY users.username;`
const getReceivedCoinsQuery = `SELECT users.username as fromUser, SUM(amount_of_coins) as amount FROM public.coin_transactions
			INNER JOIN users ON users.id = from_id
			WHERE to_id = $1
			GROUP BY users.username;`

type InfoStorage interface {
	GetInfoResponse(uuid.UUID) (api.InfoResponse, error)
}

type InfoStorageImpl struct {
	db *sqlx.DB
}
type ScanFunction[T any] func(rows *sql.Rows, dto *T) error

func scanBoughts(rows *sql.Rows, merchDto *api.MerchDto) error {
	return rows.Scan(&merchDto.Type, &merchDto.Quanitity)
}

func scanSend(rows *sql.Rows, toUserDto *api.ToUserDto) error {
	return rows.Scan(&toUserDto.ToUser, &toUserDto.Amount)
}

func scanCoinHistory(rows *sql.Rows, fromUserDto *api.FromUserDto) error {
	return rows.Scan(&fromUserDto.FromUser, &fromUserDto.Amount)
}

func QueryHelper[T any](tx *sql.Tx, userId uuid.UUID, q string, scanF ScanFunction[T]) ([]T, error) {
	var res []T
	var dto T
	rows, err := tx.Query(q, userId)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		err := scanF(rows, &dto)
		if err != nil {
			return res, err
		}
		res = append(res, dto)
	}
	return res, err
}

func (gi *InfoStorageImpl) GetInfoResponse(userId uuid.UUID) (api.InfoResponse, error) {
	infoResponse := api.InfoResponse{}
	tx, err := gi.db.Begin()
	if err != nil {
		return infoResponse, err
	}
	boughts, err := QueryHelper(tx, userId, getBoughtsQuery, scanBoughts)
	if err != nil {
		return infoResponse, errors.Join(err, tx.Rollback())
	}
	infoResponse.Inventory = boughts
	sent, err := QueryHelper(tx, userId, getSentCoinsQuery, scanSend)
	if err != nil {
		return infoResponse, errors.Join(err, tx.Rollback())
	}
	infoResponse.Sent = sent
	coinHistory, err := QueryHelper(tx, userId, getReceivedCoinsQuery, scanCoinHistory)
	if err != nil {
		return infoResponse, errors.Join(err, tx.Rollback())
	}
	infoResponse.CoinHistory = coinHistory
	return infoResponse, tx.Commit()
}

func NewInfoStorage(db *sqlx.DB) InfoStorage {
	return &InfoStorageImpl{
		db: db,
	}
}
