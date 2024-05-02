package sqlite

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Sqlite struct {
	db *sql.DB
}

func New(url string) (*Sqlite, func() error) {
	db, err := sql.Open("sqlite", url)
	if err != nil {
		log.Fatal(err)
	}

	// Create initial DB.
	_, err = db.Exec(SchemaSqlite)
	if err != nil {
		log.Fatal(err)
	}

	return &Sqlite{
		db: db,
	}, db.Close
}
