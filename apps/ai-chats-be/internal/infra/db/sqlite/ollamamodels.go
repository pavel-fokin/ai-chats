package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"ai-chats/internal/domain"
)

type OllamaModels struct {
	db DB
}

func NewOllamaModels(db *sql.DB) *OllamaModels {
	return &OllamaModels{DB{db: db}}
}

func (m *OllamaModels) Add(ctx context.Context, model domain.OllamaModel) error {
	_, err := m.db.DBTX(ctx).ExecContext(
		ctx,
		`INSERT INTO ollama_model
		(model, added_at, updated_at, status)
		VALUES (?, ?, ?, ?)`,
		model.Model,
		model.AddedAt.Format(time.RFC3339Nano),
		model.UpdatedAt.Format(time.RFC3339Nano),
		model.Status,
	)
	return err
}

func (m *OllamaModels) AllAvailable(ctx context.Context) ([]domain.OllamaModel, error) {
	rows, err := m.db.DBTX(ctx).QueryContext(
		ctx,
		"SELECT model, added_at, updated_at, status FROM ollama_model WHERE status = ?",
		domain.StatusAvailable,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []domain.OllamaModel
	for rows.Next() {
		var (
			model              domain.OllamaModel
			addedAt, updatedAt string
		)
		err = rows.Scan(&model.Model, &addedAt, &updatedAt, &model.Status)
		if err != nil {
			return nil, err
		}

		model.AddedAt, err = time.Parse(time.RFC3339Nano, addedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse ollama_model.added_at: %w", err)
		}

		model.UpdatedAt, err = time.Parse(time.RFC3339Nano, updatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse ollama_model.updated_at: %w", err)
		}

		models = append(models, model)
	}

	return models, nil
}

func (m *OllamaModels) Delete(ctx context.Context, model domain.OllamaModel) error {
	_, err := m.db.DBTX(ctx).ExecContext(
		ctx,
		"DELETE FROM ollama_model WHERE model = ?",
		model.Model,
	)
	return err
}

func (m *OllamaModels) Exists(ctx context.Context, model string) (bool, error) {
	var exists bool
	err := m.db.DBTX(ctx).QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM ollama_model WHERE model = ?)",
		model,
	).Scan(&exists)
	return exists, err
}

func (m *OllamaModels) Save(ctx context.Context, model domain.OllamaModel) error {
	_, err := m.db.DBTX(ctx).ExecContext(
		ctx,
		`UPDATE ollama_model
		SET updated_at = ?, status = ?
		WHERE model = ?`,
		model.UpdatedAt.Format(time.RFC3339Nano),
		model.Status,
		model.Model,
	)
	return err
}
