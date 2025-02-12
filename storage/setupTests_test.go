package storage

import (
	"testing"

	"github.com/fridrock/avito-shop/testdbsetup"
	"github.com/jmoiron/sqlx"
)

var (
	userST    UserStorage
	coinST    CoinStorage
	productST ProductStorage
	conn      *sqlx.DB
)

func TestMain(m *testing.M) {
	conn = testdbsetup.CreateTestConnection("..")
	defer conn.Close()
	userST = NewUserStorage(conn)
	coinST = NewCoinStorage(conn)
	productST = NewProductStorage(conn)
	m.Run()
}
