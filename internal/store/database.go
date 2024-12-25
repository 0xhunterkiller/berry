package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() (*sqlx.DB, error) {

	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_USER"),
		os.Getenv("PSQL_PASSWORD"),
		os.Getenv("PSQL_DB"),
		os.Getenv("PSQL_SSLMODE"))

	var err error
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return db, err
}

func MigUp(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to initialize migration driver:  %w", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(os.Getenv("MIG_DIR"), "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to initialize migration instance:  %w", err)
	}

	err = mig.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("migration error: %w", err)
		} else {
			logger.Logger.Info("no changes")
		}
	}
	logger.Logger.Info("migrations up")

	return nil
}

func MigDown(db *sql.DB) error {

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to initialize migration driver: %w", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(os.Getenv("MIG_DIR"), "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to initialize migration instance: %w", err)
	}

	err = mig.Down()
	if err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("migration error: %w", err)
		} else {
			logger.Logger.Info("no changes")
		}
	}
	logger.Logger.Info("migrations down")

	return nil
}

func CloseDBConn(db *sql.DB) {
	db.Close()
}
