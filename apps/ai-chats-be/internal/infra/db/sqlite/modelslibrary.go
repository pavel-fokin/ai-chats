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

func (m *ModelsLibrary) FindDescription(ctx context.Context, modelName string) (string, error) {
	var description string
	err := m.DB.db.QueryRow("SELECT description FROM ollama_model_description WHERE model_name = ?", modelName).Scan(&description)
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
		SELECT ollama_model_description.model_name, description, GROUP_CONCAT(tag) AS tags
		FROM ollama_model_description
		LEFT JOIN ollama_model_tag ON ollama_model_description.model_name = ollama_model_tag.model_name
		GROUP BY ollama_model_description.model_name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var modelCard domain.ModelCard
		err := rows.Scan(&modelCard.ModelName, &modelCard.Description, &tags)
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

func (m *ModelsLibrary) FindByName(ctx context.Context, modelName string) (*domain.ModelCard, error) {
	var (
		modelCard domain.ModelCard
		tags      string
	)
	err := m.DB.db.QueryRow(`
		SELECT ollama_model_description.model_name, description, GROUP_CONCAT(tag) AS tags
		FROM ollama_model_description
		LEFT JOIN ollama_model_tag ON ollama_model_description.model_name = ollama_model_tag.model_name
		WHERE ollama_model_description.model_name = ?
		GROUP BY ollama_model_description.model_name`, modelName).Scan(&modelCard.ModelName, &modelCard.Description, &tags)
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
