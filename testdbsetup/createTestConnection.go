package testdbsetup

import (
	"context"
	"log"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

func CreateTestConnection(prefix string) *sqlx.DB {
	ctx := context.Background()
	connString, err := GetDatabaseContainer(prefix).ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatal("error creating connection string" + err.Error())
	}
	connection, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatal("error opening connection" + err.Error())
	}
	slog.Info("successful creating of test container")
	return connection
}
