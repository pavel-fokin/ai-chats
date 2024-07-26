package sqlite

import (
	"context"
	"testing"

	"ai-chats/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestSqliteOllamaModels(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	ollamaModels := NewOllamaModels(db)

	t.Run("all added is empty", func(t *testing.T) {
		models, err := ollamaModels.AllAdded(ctx)
		assert.NoError(t, err)
		assert.Empty(t, models)
	})

	t.Run("valid", func(t *testing.T) {
		err := ollamaModels.Add(ctx, *domain.NewOllamaModel("model"))
		assert.NoError(t, err)
	})

	t.Run("add model with empty model", func(t *testing.T) {
		err := ollamaModels.Add(ctx, *domain.NewOllamaModel(""))
		assert.Error(t, err)
	})

	t.Run("delete model", func(t *testing.T) {
		model := domain.NewOllamaModel("model")
		model.Delete()
		err := ollamaModels.Delete(ctx, *model)
		assert.NoError(t, err)
	})

	t.Run("all added", func(t *testing.T) {
		db := New(":memory:")
		defer db.Close()
		CreateTables(db)

		ollamaModels := NewOllamaModels(db)

		_, err := db.Exec(
			`INSERT INTO ollama_model
			(model, added_at, updated_at, status)
			VALUES ('model1', '2006-01-02T15:04:05.999999999Z', '2006-01-02T15:04:05.999999999Z', 'added')`)
		assert.NoError(t, err)
		_, err = db.Exec(
			`INSERT INTO ollama_model
			(model, added_at, updated_at, status)
			VALUES ('model2', '2006-01-02T15:04:05.999999999Z', '2006-01-02T15:04:05.999999999Z', 'added')`)
		assert.NoError(t, err)

		models, err := ollamaModels.AllAdded(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, models)
		assert.Len(t, models, 2)
	})
}
