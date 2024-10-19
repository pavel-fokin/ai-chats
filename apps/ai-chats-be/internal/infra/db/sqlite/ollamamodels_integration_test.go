package sqlite

import (
	"ai-chats/internal/domain"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOllamaModels_FindOllamaModelsPullingInProgress(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	ollamaModels := NewOllamaModels(db)

	t.Run("no models", func(t *testing.T) {
		models, err := ollamaModels.FindOllamaModelsPullingInProgress(ctx)
		assert.NoError(t, err)
		assert.Empty(t, models)
	})

	t.Run("some models", func(t *testing.T) {
		ollamaModels.AddModelPullingStarted(ctx, "model1")
		models, err := ollamaModels.FindOllamaModelsPullingInProgress(ctx)
		assert.NoError(t, err)
		assert.Equal(t, []domain.OllamaModel{
			{
				Model:  "model1",
				Status: domain.OllamaModelStatusPulling,
			},
		}, models)
	})

}
