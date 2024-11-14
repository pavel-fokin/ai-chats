package sqlite

import (
	"context"
	"database/sql"
	"strings"

	"ai-chats/internal/domain"
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

func (m *ModelsLibrary) FindAll(ctx context.Context) ([]*domain.ModelCard, error) {
	var (
		modelCards []*domain.ModelCard
		tags       string
	)

	rows, err := m.DB.db.Query(`
		SELECT name, description, GROUP_CONCAT(tag) AS tags
		FROM model_description
		LEFT JOIN model_tag ON model_description.name = model_tag.model
		GROUP BY model_description.name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var modelCard domain.ModelCard
		err := rows.Scan(&modelCard.Model, &modelCard.Description, &tags)
		if err != nil {
			return nil, err
		}
		modelCard.Tags = strings.Split(tags, ",")
		modelCards = append(modelCards, &modelCard)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return modelCards, nil
}

func (m *ModelsLibrary) FindByName(ctx context.Context, name string) (*domain.ModelCard, error) {
	var (
		modelCard domain.ModelCard
		tags      string
	)
	err := m.DB.db.QueryRow(`
		SELECT name, description, GROUP_CONCAT(tag) AS tags
		FROM model_description
		LEFT JOIN model_tag ON model_description.name = model_tag.model
		WHERE model_description.name = ?
		GROUP BY model_description.name`, name).Scan(&modelCard.Model, &modelCard.Description, &tags)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, domain.ErrModelNotFound
		default:
			return nil, err
		}
	}

	modelCard.Tags = strings.Split(tags, ",")

	return &modelCard, nil
}
