package sqlite

import (
	"context"
	"database/sql"
	"time"

	"ai-chats/internal/domain"
)

type OllamaModels struct {
	DB
}

func NewOllamaModels(db *sql.DB) *OllamaModels {
	return &OllamaModels{DB{db: db}}
}

func (m *OllamaModels) AddModelPullingStarted(ctx context.Context, model string) error {
	startedAt := time.Now().UTC().Format(time.RFC3339Nano)

	_, err := m.DB.db.Exec("INSERT INTO ollama_model_pulling (model, started_at) VALUES (?, ?)", model, startedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *OllamaModels) AddModelPullingFinished(ctx context.Context, model string, finalStatus domain.OllamaPullingFinalStatus) error {
	finishedAt := time.Now().UTC().Format(time.RFC3339Nano)

	_, err := m.DB.db.Exec("UPDATE ollama_model_pulling SET finished_at = ?, final_status = ? WHERE model = ?", finishedAt, finalStatus, model)
	if err != nil {
		return err
	}

	return nil
}

func (m *OllamaModels) FindOllamaModelsPullingInProgress(ctx context.Context) ([]domain.OllamaModel, error) {
	rows, err := m.DB.db.Query("SELECT model FROM ollama_model_pulling WHERE finished_at IS NULL")
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
