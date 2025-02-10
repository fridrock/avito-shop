package buy

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BuyStorage interface {
	Buy(uuid.UUID, string) error
}

type BuyStorageImpl struct {
	db *sqlx.DB
}

func (bs *BuyStorageImpl) Buy(userId uuid.UUID, itemName string) error {
	//find item id

	// q := `INSERT INTO boughts(id, hashed_password, coins) VALUES ($1, $2, $3, 1000) RETURNING id`
	// var id uuid.UUID
	// err := bs.db.QueryRow(
	// 	q,
	// 	uuid.New().String(),
	// 	user.Username,
	// 	user.HashedPassword).Scan(&id)
	return nil
}

func NewBuyStorage(db *sqlx.DB) BuyStorage {
	return &BuyStorageImpl{
		db: db,
	}
}
