package sqlite

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func New(url string) *sql.DB {
	db, err := sql.Open("sqlite", url)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	return db
}

func CreateTables(db *sql.DB) {
	_, err := db.Exec(SchemaSqlite)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}

func LoadFixtures(db *sql.DB) {
	_, err := db.Exec(FixturesSqlite)
	if err != nil {
		log.Fatalf("Failed to load fixtures: %v", err)
	}
}

// DBTX is an interface for database connections or transactions.
type DBTX interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// DB is a database connection.
type DB struct {
	db *sql.DB
}

// DBTX returns a transaction from the context if it exists, otherwise returns the database.
func (e *DB) DBTX(ctx context.Context) DBTX {
	tx := MaybeHaveTx(ctx)
	if tx != nil {
		return tx
	}

	return e.db
}
