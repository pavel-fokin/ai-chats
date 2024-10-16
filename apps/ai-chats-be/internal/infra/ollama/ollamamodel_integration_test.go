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
		modelName := "smollm:135m"

		llm, err := ollamaClient.NewModel(domain.NewOllamaModel(modelName))
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
		modelName := "model-not-found"
		llm, err := ollamaClient.NewModel(domain.NewOllamaModel(modelName))
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
