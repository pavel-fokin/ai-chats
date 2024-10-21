package sqlite

import (
	"context"
	"database/sql"

	"ai-chats/internal/domain"
)

type OllamaModels struct {
	DB
}

var _ domain.OllamaModels = (*OllamaModels)(nil)

func NewOllamaModels(db *sql.DB) *OllamaModels {
	return &OllamaModels{DB{db: db}}
}

func (m *OllamaModels) Save(ctx context.Context, model domain.OllamaModel) error {
	for _, event := range model.Events {
		if err := m.saveEvent(ctx, event); err != nil {
			return err
		}
	}

	return nil
}

func (m *OllamaModels) FindOllamaModelsPullInProgress(ctx context.Context) ([]domain.OllamaModel, error) {
	rows, err := m.DB.db.Query(
		`SELECT model
		FROM (
			SELECT model, type, MAX(occurred_at) AS last_occurred_at
			FROM ollama_model_pull_event
			GROUP BY model
		) AS last_events
		WHERE type = ?;`,
		string(domain.OllamaModelPullStartedType),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []domain.OllamaModel
	for rows.Next() {
		var model string
		if err := rows.Scan(&model); err != nil {
			return nil, err
		}

		ollamaModel, _ := domain.NewOllamaModel(model)
		ollamaModel.SetStatus(domain.OllamaModelStatusPulling)
		models = append(models, ollamaModel)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return models, nil
}

// Save OllamaModelPullEvent to the database.
func (m *OllamaModels) saveEvent(ctx context.Context, event domain.OllamaModelPullEvent) error {
	_, err := m.DBTX(ctx).Exec(
		"INSERT INTO ollama_model_pull_event (id, model, occurred_at, type) VALUES (?, ?, ?, ?)",
		event.ID(), event.Model(), event.OccurredAt(), event.Type(),
	)
	if err != nil {
		return err
	}

	return nil
}
