package storage

import (
	"context"
	"log"
	"log/slog"
	"testing"

	"github.com/jmoiron/sqlx"
)

var (
	userStorage    UserStorage
	coinStorage    CoinStorage
	productStorage ProductStorage
	conn           *sqlx.DB
)

func TestMain(m *testing.M) {
	createConnection()
	defer conn.Close()
	userStorage = NewUserStorage(conn)
	coinStorage = NewCoinStorage(conn)
	productStorage = NewProductStorage(conn)
	m.Run()
}

func createConnection() {
	ctx := context.Background()
	connString, err := GetDatabaseContainer().ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatal("error creating connection string" + err.Error())
	}
	connection, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatal("error opening connection" + err.Error())
	}
	slog.Info("successful creating of test container")
	conn = connection
}
