package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/commands"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain/events"
	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/llm"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/json"
)

// GenerateResponse generates a LLM response for the chat.
func (a *App) GenerateResponse(ctx context.Context, chatID uuid.UUID) error {
	messages, err := a.AllMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to get messages: %w", err)
	}

	llm, err := llm.NewOllama("llama3")
	if err != nil {
		return fmt.Errorf("failed to create a chat model: %w", err)
	}

	streamFunc := func(messageChunk events.MessageChunkReceived) error {
		messageChunkReceived := events.NewMessageChunkReceived(
			messageChunk.MessageID,
			messageChunk.Sender,
			messageChunk.Text,
			messageChunk.Final,
		)
		if err := a.pubsub.Publish(ctx, chatID.String(), json.MustMarshal(ctx, messageChunkReceived)); err != nil {
			return fmt.Errorf("failed to publish a message chunk received event: %w", err)
		}
		return nil
	}

	llmMessage, err := llm.GenerateResponseWithStream(ctx, messages, streamFunc)
	if err != nil {
		return fmt.Errorf("failed to generate a response: %w", err)
	}

	if err := a.messages.Add(ctx, chatID, llmMessage); err != nil {
		return fmt.Errorf("failed to add a message: %w", err)
	}

	messageSent := events.NewMessageAdded(chatID, llmMessage)
	if err := a.pubsub.Publish(ctx, chatID.String(), json.MustMarshal(ctx, messageSent)); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	if len(messages) == 1 {
		generateTitle := commands.NewGenerateTitle(chatID)
		if err := a.pubsub.Publish(ctx, "generate-title", json.MustMarshal(ctx, generateTitle)); err != nil {
			return fmt.Errorf("failed to publish a generate title command: %w", err)
		}
	}

	return nil
}

// GenerateTitle generates a LLM title for the chat.
func (a *App) GenerateTitle(ctx context.Context, chatID uuid.UUID) error {
	messages, err := a.AllMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to get messages: %w", err)
	}

	llm, err := llm.NewOllama("llama3")
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
