//go:build integration

package ollama

import (
	"context"
	"ai-chats/internal/domain"
	"ai-chats/internal/domain/events"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOllamaGenerateResponse(t *testing.T) {
	user := domain.NewUser("username", "password")

	llm, err := NewOllama(domain.NewOllamaModel("llama3"))
	assert.NoError(t, err)

	llmResponse, err := llm.GenerateResponse(context.Background(), []domain.Message{
		{Sender: domain.NewUserSender(user), Text: "Hi"},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, llmResponse.Text)
}

func TestOllamaGenerateResponseWithStream(t *testing.T) {
	user := domain.NewUser("username", "password")

	llm, err := NewOllama(domain.NewOllamaModel("llama3"))
	assert.NoError(t, err)

	messageChunkReceived := events.MessageChunkReceived{}
	llmResponse, err := llm.GenerateResponseWithStream(
		context.Background(),
		[]domain.Message{
			{Sender: domain.NewUserSender(user), Text: "Hi"},
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
