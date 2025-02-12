package storage

import (
	"testing"

	"github.com/fridrock/avito-shop/testdbsetup"
	"github.com/jmoiron/sqlx"
)

var (
	userStorage    UserStorage
	coinStorage    CoinStorage
	productStorage ProductStorage
	conn           *sqlx.DB
)

func TestMain(m *testing.M) {
	conn = testdbsetup.CreateTestConnection("..")
	defer conn.Close()
	userStorage = NewUserStorage(conn)
	coinStorage = NewCoinStorage(conn)
	productStorage = NewProductStorage(conn)
	m.Run()
}
