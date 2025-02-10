package storage

import (
	"github.com/fridrock/avito-shop/api"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type GetInfoStorage interface {
	GetInfo()
}

type GetInfoStorageImpl struct {
	db *sqlx.DB
}

func (gi *GetInfoStorageImpl) GetInfoResponse(userId uuid.UUID) (api.InfoResponse, error) {
	infoResponse := api.InfoResponse{}
	q := `SELECT products.product_name as type, COUNT(*) as quantity FROM public.boughts
            INNER JOIN products ON products.id = product_id
            WHERE user_id = $1
            GROUP BY products.product_name;`
	err := gi.db.Select(&infoResponse.Inventory, q, userId)
	if err != nil {
		return infoResponse, err
	}
	return infoResponse, nil
}
