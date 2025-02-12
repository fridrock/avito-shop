package testdbsetup

import (
	"context"
	"log"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	pgContainer *postgres.PostgresContainer
)

func initPostgresContainer(prefix string) {
	ctx := context.Background()

	dbName := "shop"
	dbUser := "postgres"
	dbPassword := "password"

	containerCreated, err := postgres.Run(ctx,
		"postgres:latest",
		postgres.WithInitScripts(filepath.Join(prefix, "migrations", "testInit.sql")),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}
	pgContainer = containerCreated
}

func GetDatabaseContainer(prefix string) *postgres.PostgresContainer {
	if pgContainer == nil {
		initPostgresContainer(prefix)
	}
	return pgContainer
}

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
