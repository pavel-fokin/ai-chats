package sqlite

import "database/sql"

type OllamaModels struct {
	db DB
}

func NewOllamaModels(db *sql.DB) *OllamaModels {
	return &OllamaModels{DB{db: db}}
}
