package storage

import (
	"context"
	"log"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	pgContainer *postgres.PostgresContainer
)

func initPostgresContainer() {
	ctx := context.Background()

	dbName := "shop"
	dbUser := "postgres"
	dbPassword := "password"

	containerCreated, err := postgres.Run(ctx,
		"postgres:latest",
		postgres.WithInitScripts(filepath.Join("./", "testInit.sql")),
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

func GetDatabaseContainer() *postgres.PostgresContainer {
	if pgContainer == nil {
		initPostgresContainer()
	}
	return pgContainer
}
