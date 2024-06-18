//go:build integration

package llm

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOllamaModels(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	ollama, err := NewOllamaModels()
	assert.NoError(err)

	t.Run("All", func(t *testing.T) {
		_, err := ollama.All(ctx)
		assert.NoError(err)
	})

	t.Run("Pull", func(t *testing.T) {
		err := ollama.Pull(ctx, domain.NewModel("all-minilm", "latest"))
		assert.NoError(err)
	})

	t.Run("Check if model exists", func(t *testing.T) {
		models, err := ollama.All(ctx)
		assert.NoError(err)
		assert.NotEmpty(models)

		modelExists := false
		for _, model := range models {
			if model.Name == "all-minilm" && model.Tag == "latest" {
				modelExists = true
				break
			}
		}

		assert.True(modelExists)
	})

	t.Run("Delete", func(t *testing.T) {
		err := ollama.Delete(ctx, domain.NewModel("all-minilm", "latest"))
		assert.NoError(err)
	})
}
