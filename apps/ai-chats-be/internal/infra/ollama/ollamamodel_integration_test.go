//go:build integration

package ollama

import (
	"ai-chats/internal/domain"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOllamaModel_Chat(t *testing.T) {
	t.Parallel()
	ollamaClient := NewOllamaClient()

	t.Run("ollama chat", func(t *testing.T) {
		model := "smollm:135m"

		ollamaModel, err := domain.NewOllamaModel(model)
		assert.NoError(t, err)

		llm, err := ollamaClient.NewModel(ollamaModel)
		assert.NoError(t, err)

		llmMessage, err := llm.Chat(
			context.Background(),
			[]domain.Message{
				{Sender: domain.NewUserSender(domain.NewUserID()), Text: "Hi"},
			},
			func(msg domain.Message) error {
				return nil
			},
		)

		assert.NoError(t, err)
		assert.NotEmpty(t, llmMessage.Text)
	})

	t.Run("model not found", func(t *testing.T) {
		model := "model-not-found"
		ollamaModel, err := domain.NewOllamaModel(model)
		assert.NoError(t, err)

		llm, err := ollamaClient.NewModel(ollamaModel)
		assert.NoError(t, err)

		llmMessage, err := llm.Chat(
			context.Background(),
			[]domain.Message{
				{Sender: domain.NewUserSender(domain.NewUserID()), Text: "Hi"},
			},
			func(msg domain.Message) error {
				return nil
			},
		)

		assert.Error(t, err)
		assert.Empty(t, llmMessage.Text)
	})
}
