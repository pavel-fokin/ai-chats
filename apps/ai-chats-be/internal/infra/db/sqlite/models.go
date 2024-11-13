package sqlite

import (
	"context"
	"database/sql"
)

type ModelsLibrary struct {
	DB
}

func NewModelsLibrary(db *sql.DB) *ModelsLibrary {
	return &ModelsLibrary{DB{db: db}}
}

func (m *ModelsLibrary) FindDescription(ctx context.Context, name string) (string, error) {
	var description string
	err := m.DB.db.QueryRow("SELECT description FROM model_description WHERE name = ?", name).Scan(&description)
	if err != nil {
		return "", err
	}

	return description, nil
}
