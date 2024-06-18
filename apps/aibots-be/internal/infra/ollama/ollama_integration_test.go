//go:build integration

package ollama

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain/events"
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

func TestOllamaGenerateResponseWithStream(t *testing.T) {
	llm, err := NewOllama("llama3")
	assert.NoError(t, err)

	messageChunkReceived := events.MessageChunkReceived{}
	llmResponse, err := llm.GenerateResponseWithStream(
		context.Background(),
		[]domain.Message{
			{Sender: "User", Text: "Hi"},
		},
		func(messageChunk events.MessageChunkReceived) error {
			messageChunkReceived = messageChunk
			return nil
		},
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, llmResponse.Text)
	assert.True(t, messageChunkReceived.Final)
}
