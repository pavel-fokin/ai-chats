package app

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain/events"
	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/ollama"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/json"
	"pavel-fokin/ai/apps/ai-bots-be/internal/worker"
)

// GenerateResponse generates a LLM response for the chat.
func (a *App) GenerateResponse(ctx context.Context, chatID domain.ChatID) error {
	chat, err := a.chats.FindByID(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	messages, err := a.chats.AllMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to get messages: %w", err)
	}
	llm, err := ollama.NewOllama(chat.DefaultModel.AsOllamaModel())
	if err != nil {
		return fmt.Errorf("failed to create a chat model: %w", err)
	}

	streamFunc := func(messageChunk events.MessageChunkReceived) error {
		if err := a.pubsub.Publish(ctx, chatID.String(), json.MustMarshal(ctx, messageChunk)); err != nil {
			return fmt.Errorf("failed to publish a message chunk received event: %w", err)
		}
		return nil
	}

	llmMessage, err := llm.GenerateResponseWithStream(ctx, messages, streamFunc)
	if err != nil {
		return fmt.Errorf("failed to generate a response: %w", err)
	}

	if err := a.chats.AddMessage(ctx, chatID, llmMessage); err != nil {
		return fmt.Errorf("failed to add a message: %w", err)
	}

	messageAdded := events.NewMessageAdded(chatID, llmMessage)
	if err := a.pubsub.Publish(ctx, worker.MessageAddedTopic, json.MustMarshal(ctx, messageAdded)); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	return nil
}

// GenerateTitle generates a LLM title for the chat.
func (a *App) GenerateTitle(ctx context.Context, chatID domain.ChatID) error {
	chat, err := a.chats.FindByID(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	messages, err := a.AllMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to get messages: %w", err)
	}

	llm, err := ollama.NewOllama(chat.DefaultModel.AsOllamaModel())
	if err != nil {
		return fmt.Errorf("failed to create a chat model: %w", err)
	}

	generatedTitle, err := llm.GenerateTitle(ctx, messages)
	if err != nil {
		return fmt.Errorf("failed to generate a title: %w", err)
	}

	if err := a.chats.UpdateTitle(ctx, chatID, generatedTitle); err != nil {
		return fmt.Errorf("failed to update chat title: %w", err)
	}

	// titleUpdated := events.NewChatTitleUpdated(chatID, generatedTitle)
	// if err := a.events.Publish(ctx, chatID.String(), json.MustMarshal(ctx, titleUpdated)); err != nil {
	// 	return fmt.Errorf("failed to publish a title updated event: %w", err)
	// }

	return nil
}
