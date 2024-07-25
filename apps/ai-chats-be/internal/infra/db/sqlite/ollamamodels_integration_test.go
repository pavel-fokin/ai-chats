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

	t.Run("valid", func(t *testing.T) {
		err := ollamaModels.Add(ctx, *domain.NewOllamaModel("model"))
		assert.NoError(t, err)
	})

	t.Run("add model with empty model", func(t *testing.T) {
		err := ollamaModels.Add(ctx, *domain.NewOllamaModel(""))
		assert.Error(t, err)
	})

	t.Run("all available is empty", func(t *testing.T) {
		models, err := ollamaModels.AllAvailable(ctx)
		assert.NoError(t, err)
		assert.Empty(t, models)
	})

	t.Run("delete model", func(t *testing.T) {
		err := ollamaModels.Delete(ctx, *domain.NewOllamaModel("model"))
		assert.NoError(t, err)
	})

	t.Run("all available", func(t *testing.T) {
		db.Exec(
			`INSERT INTO ollama_model
		(model, added_at, updated_at, status)
		VALUES ('model1', '2006-01-02T15:04:05.999999999Z', '2006-01-02T15:04:05.999999999Z', 'available')`)
		db.Exec(
			`INSERT INTO ollama_model
		(model, added_at, updated_at, status)
		VALUES ('model2', '2006-01-02T15:04:05.999999999Z', '2006-01-02T15:04:05.999999999Z', 'available')`)

		models, err := ollamaModels.AllAvailable(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, models)
		assert.Len(t, models, 2)
	})
}
