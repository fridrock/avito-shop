package storage

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	Id    int    `db:"id"`
	Name  string `db:"product_name"`
	Price int    `db:"price"`
}

type ProductStorage interface {
	Buy(uuid.UUID, Product) error
	FindProductByName(string) (Product, error)
}

type ProductStorageImpl struct {
	db *sqlx.DB
}

func (bs *ProductStorageImpl) Buy(userId uuid.UUID, product Product) error {
	tx, err := bs.db.Begin()
	if err != nil {
		return err
	}
	q := `UPDATE users SET coins = coins - $1 WHERE id = $2`
	_, err1 := bs.db.Exec(q, product.Price, userId.String())
	q = `INSERT INTO boughts(bought_id, user_id, product_id) VALUES ($1, $2, $3)`
	_, err2 := bs.db.Exec(
		q,
		uuid.New().String(),
		userId,
		product.Id)
	resErr := errors.Join(err1, err2)
	if resErr != nil {
		rollbackErr := tx.Rollback()
		return errors.Join(resErr, rollbackErr)
	}
	return tx.Commit()
}

func (bs *ProductStorageImpl) FindProductByName(itemName string) (Product, error) {
	q := `SELECT * FROM products WHERE product_name = $1`
	var product Product
	err := bs.db.Get(&product, q, itemName)
	return product, err
}

func NewProductStorage(db *sqlx.DB) ProductStorage {
	return &ProductStorageImpl{
		db: db,
	}
}
