//go:build integration

package llm

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOllamaGenerateResponse(t *testing.T) {
	llm, err := NewOllama("llama3")
	assert.NoError(t, err)

	llmResponse, err := llm.GenerateResponse(context.Background(), []domain.Message{
		{Sender: "User", Text: "Hi"},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, llmResponse.Text)
}
