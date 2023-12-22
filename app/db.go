package app

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	PgxMigration "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"projnellis.com/menhir/db"
)

type Database struct {
	Context  context.Context
	Queries  *db.Queries
	Database *pgxpool.Pool
}

func (database *Database) Init(dsn string) {
	ctx := context.Background()
	cfg, errConfig := pgxpool.ParseConfig(dsn)
	if errConfig != nil {
		log.Fatal("PGXPOOL failed to parse the dsn.")
		os.Exit(-1)
	}

	dbConn, errConnectConfig := pgxpool.ConnectConfig(ctx, cfg)
	if errConnectConfig != nil {
		log.Fatalf("Failed to connect to database: %v", errConnectConfig)
		os.Exit(-1)
	}
	queries := db.New(dbConn)
	database.Database = dbConn
	database.Context = ctx
	database.Queries = queries
}

func RunMigrate(dsn string) {
	instance, errOpen := sql.Open("pgx", dsn)
	if errOpen != nil {
		log.Fatal("Failed to open database for migration")
		os.Exit(-1)
	}
	if errPing := instance.Ping(); errPing != nil {
		log.Fatal("Cannot migrate, failed to connect to target server")
		os.Exit(-1)
	}

	driver, _ := PgxMigration.WithInstance(instance, &PgxMigration.Config{
		MigrationsTable:       "_migration",
		SchemaName:            "public",
		StatementTimeout:      60 * time.Second,
		MultiStatementEnabled: true,
	})

	migrate, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"pgx", driver,
	)

	if err != nil {
		log.Fatal("Failed to create migration")
		os.Exit(-1)
	}
	err = migrate.Up()
	if err != nil {
		log.Println("Failed to run migrations: ", err)
	}
}
