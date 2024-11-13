//go:build integration

package ollama

import (
	"ai-chats/internal/domain"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOllamaModels(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	assert := assert.New(t)

	ollama := NewOllamaClient()
	modelName := "all-minilm"

	t.Run("list", func(t *testing.T) {
		_, err := ollama.List(ctx)
		assert.NoError(err)
	})

	t.Run("pull", func(t *testing.T) {
		err := ollama.Pull(ctx, modelName, func(progress domain.OllamaModelPullProgress) error {
			return nil
		})
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
		err := ollama.Delete(ctx, modelName)
		assert.NoError(err)
	})
}
