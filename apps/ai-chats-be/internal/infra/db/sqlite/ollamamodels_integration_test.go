package sqlite

import (
	"ai-chats/internal/domain"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOllamaModels_Save(t *testing.T) {
	ctx := context.Background()

	db := New(":memory:")
	defer db.Close()
	CreateTables(db)

	ollamaModels := NewOllamaModels(db)

	t.Run("save model with no events", func(t *testing.T) {
		model, err := domain.NewOllamaModel("model1")
		assert.NoError(t, err)

		err = ollamaModels.Save(ctx, model)
		assert.NoError(t, err)
	})

	t.Run("save model with events", func(t *testing.T) {
		model, err := domain.NewOllamaModel("model2")
		assert.NoError(t, err)

		model.PullStarted()
		model.PullCompleted()
		model.PullFailed()
		err = ollamaModels.Save(ctx, model)
		assert.NoError(t, err)
	})
}

func TestOllamaModels_FindOllamaModelPullInProgress(t *testing.T) {
	ctx := context.Background()

	t.Run("no models", func(t *testing.T) {
		db := New(":memory:")
		defer db.Close()
		CreateTables(db)

		ollamaModels := NewOllamaModels(db)

		models, err := ollamaModels.FindOllamaModelsPullInProgress(ctx)
		assert.NoError(t, err)
		assert.Empty(t, models)
	})

	t.Run("some models", func(t *testing.T) {
		db := New(":memory:")
		defer db.Close()
		CreateTables(db)

		ollamaModels := NewOllamaModels(db)

		modelInProgress, _ := domain.NewOllamaModel("model1")
		modelInProgress.PullStarted()
		ollamaModels.Save(ctx, modelInProgress)

		modelCompleted, _ := domain.NewOllamaModel("model2")
		modelCompleted.PullStarted()
		modelCompleted.PullCompleted()
		ollamaModels.Save(ctx, modelCompleted)

		modelFailed, _ := domain.NewOllamaModel("model3")
		modelFailed.PullStarted()
		modelFailed.PullFailed()
		ollamaModels.Save(ctx, modelFailed)

		models, err := ollamaModels.FindOllamaModelsPullInProgress(ctx)
		assert.NoError(t, err)
		assert.Equal(t, []domain.OllamaModel{
			{
				Model:  "model1",
				Name:   "model1",
				Tag:    "latest",
				Status: domain.OllamaModelStatusPulling,
			},
		}, models)
	})

	t.Run("pull started 3 times", func(t *testing.T) {
		db := New(":memory:")
		defer db.Close()
		CreateTables(db)

		ollamaModels := NewOllamaModels(db)

		model, err := domain.NewOllamaModel("model4")
		assert.NoError(t, err)

		model.PullStarted()
		model.PullCompleted()
		model.PullStarted()
		model.PullFailed()
		model.PullStarted()
		err = ollamaModels.Save(ctx, model)
		assert.NoError(t, err)

		models, err := ollamaModels.FindOllamaModelsPullInProgress(ctx)
		assert.NoError(t, err)
		assert.Equal(t, []domain.OllamaModel{
			{
				Model:  "model4",
				Name:   "model4",
				Tag:    "latest",
				Status: domain.OllamaModelStatusPulling,
			},
		}, models)
	})
}
