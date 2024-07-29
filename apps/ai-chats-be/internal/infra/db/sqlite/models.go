package sqlite

import (
	"context"
	"database/sql"
)

type Models struct {
	DB
}

func NewModels(db *sql.DB) *Models {
	return &Models{DB{db: db}}
}

func (m *Models) FindDescription(ctx context.Context, name string) (string, error) {
	var description string
	err := m.DB.db.QueryRow("SELECT description FROM model_description WHERE name = ?", name).Scan(&description)
	if err != nil {
		return "", err
	}

	return description, nil
}
