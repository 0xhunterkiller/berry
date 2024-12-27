package dbpsql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

// ConnectDB establishes a connection to the PostgreSQL database with connection pooling configurations.
// It reads database configuration values (host, port, user, password, database name, and SSL mode) from environment variables.
// The function also allows setting the maximum number of open connections, maximum idle connections, and maximum connection lifetime.
//
// Parameters:
//   - maxOpenConnections int: The maximum number of open connections to the database.
//   - maxIdleConnections int: The maximum number of idle connections in the connection pool.
//   - connMaxLifetimeMins int: The maximum lifetime of a connection in minutes.
//
// Returns:
//   - *sqlx.DB: A pointer to the sqlx.DB object representing the database connection.
//   - error: An error object if the connection fails or configuration encounters an issue, otherwise nil.
func ConnectDB(maxOpenConnections int, maxIdleConnections int, connMaxLifetimeMins int) (*sqlx.DB, error) {

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
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	db.SetMaxOpenConns(maxOpenConnections)
	log.Println("set number of max open connections to  ", maxOpenConnections)

	db.SetMaxIdleConns(maxIdleConnections)
	log.Println("set number of max idle connections to  ", maxOpenConnections)

	db.SetConnMaxLifetime(time.Duration(connMaxLifetimeMins) * time.Minute)
	log.Printf("set connection max lifetime to  %v mins\n", connMaxLifetimeMins)

	return db, err
}

// MigUp applies all pending migrations to the PostgreSQL database.
// It uses the Golang Migrate library to handle database migrations. The migration files are sourced from
// the directory specified by the `MIG_DIR` environment variable.
//
// Parameters:
//   - db *sqlx.DB: A pointer to the sqlx.DB object representing the database connection.
//
// Returns:
//   - error: An error object if the migration process fails, otherwise nil.
func MigUp(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
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
			fmt.Println("no changes")
		}
	}
	fmt.Println("migrations up")

	return nil
}

// MigDown rolls back all applied migrations on the PostgreSQL database.
// It uses the Golang Migrate library to handle database migrations. The migration files are sourced from
// the directory specified by the `MIG_DIR` environment variable.
//
// Parameters:
//   - db *sqlx.DB: A pointer to the sqlx.DB object representing the database connection.
//
// Returns:
//   - error: An error object if the rollback process fails, otherwise nil.
func MigDown(db *sqlx.DB) error {

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
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
			fmt.Println("no changes")
		}
	}
	fmt.Println("migrations down")

	return nil
}

// CloseDBConn closes the database connection.
// This function ensures that the resources associated with the database connection
// are properly released.
//
// Parameters:
//   - db *sqlx.DB: A pointer to the sqlx.DB object representing the database connection.
func CloseDBConn(db *sqlx.DB) {
	db.Close()
}
