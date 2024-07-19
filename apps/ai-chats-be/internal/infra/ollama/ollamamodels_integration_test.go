//go:build integration

package ollama

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOllamaModels(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	ollama := NewOllamaModels()

	t.Run("list", func(t *testing.T) {
		_, err := ollama.List(ctx)
		assert.NoError(err)
	})

	t.Run("pull", func(t *testing.T) {
		err := ollama.Pull(ctx, domain.NewOllamaModel("all-minilm"))
		assert.NoError(err)
	})

	t.Run("check if model exists", func(t *testing.T) {
		models, err := ollama.List(ctx)
		assert.NoError(err)
		assert.NotEmpty(models)

		modelExists := false
		for _, model := range models {
			if model.Model == "all-minilm:latest" {
				modelExists = true
				break
			}
		}

		assert.True(modelExists)
	})

	t.Run("delete", func(t *testing.T) {
		err := ollama.Delete(ctx, domain.NewOllamaModel("all-minilm"))
		assert.NoError(err)
	})
}
