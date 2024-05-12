package sqlite

import (
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
