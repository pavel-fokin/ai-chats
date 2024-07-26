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

func (m *OllamaModels) AllAdded(ctx context.Context) ([]domain.OllamaModel, error) {
	rows, err := m.db.DBTX(ctx).QueryContext(
		ctx,
		"SELECT model, added_at, updated_at, status FROM ollama_model WHERE status = ?",
		domain.StatusAdded,
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
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over ollama_model rows: %w", err)
	}

	for i := range models {
		var description string
		err = m.db.DBTX(ctx).QueryRowContext(
			ctx,
			"SELECT description FROM model_card WHERE model = ?",
			models[i].Name(),
		).Scan(&description)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				continue
			default:
				return nil, fmt.Errorf("failed to get model card description: %w", err)
			}
		}
		models[i].Description = description
	}

	return models, nil
}

func (m *OllamaModels) Delete(ctx context.Context, model domain.OllamaModel) error {
	if model.Status != domain.StatusDeleted {
		return domain.ErrOllamaModelNotMarkedAsDelted
	}

	_, err := m.db.DBTX(ctx).ExecContext(
		ctx,
		`UPDATE ollama_model
		SET deleted_at = ?, status = ?
		WHERE model = ?`,
		model.DeletedAt.Format(time.RFC3339Nano),
		model.Status,
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

func (o *OllamaModels) Find(ctx context.Context, model string) (domain.OllamaModel, error) {
	var (
		om                 domain.OllamaModel
		addedAt, updatedAt string
	)
	err := o.db.DBTX(ctx).QueryRowContext(
		ctx,
		"SELECT model, added_at, updated_at, status FROM ollama_model WHERE model = ?",
		model,
	).Scan(&om.Model, &addedAt, &updatedAt, &om.Status)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return om, domain.ErrOllamaModelNotFound
		default:
			return om, err
		}
	}

	om.AddedAt, err = time.Parse(time.RFC3339Nano, addedAt)
	if err != nil {
		return om, fmt.Errorf("failed to parse ollama_model.added_at: %w", err)
	}

	om.UpdatedAt, err = time.Parse(time.RFC3339Nano, updatedAt)
	if err != nil {
		return om, fmt.Errorf("failed to parse ollama_model.updated_at: %w", err)
	}

	return om, nil
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
