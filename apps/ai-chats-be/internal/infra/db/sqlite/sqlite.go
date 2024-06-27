package sqlite

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewDB(url string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", url)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(SchemaSqlite)
	return err
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
	db DBTX
}

// DBTX returns a transaction from the context if it exists, otherwise returns the database.
func (e *DB) DBTX(ctx context.Context) DBTX {
	tx := MaybeHaveTx(ctx)
	if tx != nil {
		return tx
	}

	return e.db
}
