package app

import (
	"context"
	"fmt"

	"ai-chats/internal/app/commands"
	"ai-chats/internal/domain"
	"ai-chats/internal/domain/events"
	"ai-chats/internal/infra/ollama"
	"ai-chats/internal/infra/worker"
	"ai-chats/internal/pkg/json"
)

// GenerateResponse generates a LLM response for the chat.
func (a *App) GenerateResponse(ctx context.Context, chatID domain.ChatID) error {
	chat, err := a.chats.FindByIDWithMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	llm, err := ollama.NewOllama(chat.DefaultModel.AsOllamaModel())
	if err != nil {
		return fmt.Errorf("failed to create a chat model: %w", err)
	}

	streamFunc := func(messageChunk events.MessageChunkReceived) error {
		if err := a.notifyInChat(ctx, chatID.String(), json.MustMarshal(ctx, messageChunk)); err != nil {
			return fmt.Errorf("failed to notify in chat: %w", err)
		}
		return nil
	}

	llmMessage, err := llm.GenerateResponseWithStream(ctx, chat.Messages, streamFunc)
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

// GenerateChatTitle generates a chat title.
func (a *App) GenerateChatTitleAsync(ctx context.Context, chatID domain.ChatID) error {
	generateChatTitleCommand := commands.GenerateChatTitle{ChatID: chatID.String()}
	if err := a.pubsub.Publish(
		ctx,
		worker.GenerateChatTitleTopic,
		json.MustMarshal(ctx, generateChatTitleCommand),
	); err != nil {
		return fmt.Errorf("failed to publish a generate chat title command: %w", err)
	}

	return nil
}

// GenerateTitle generates a LLM title for the chat.
func (a *App) GenerateTitle(ctx context.Context, chatID domain.ChatID) error {
	chat, err := a.chats.FindByIDWithMessages(ctx, chatID)
	if err != nil {
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	llm, err := ollama.NewOllama(chat.DefaultModel.AsOllamaModel())
	if err != nil {
		return fmt.Errorf("failed to create a chat model: %w", err)
	}

	generatedTitle, err := llm.GenerateTitle(ctx, chat.Messages)
	if err != nil {
		return fmt.Errorf("failed to generate a title: %w", err)
	}

	if err := a.chats.UpdateTitle(ctx, chatID, generatedTitle); err != nil {
		return fmt.Errorf("failed to update chat title: %w", err)
	}

	titleUpdated := events.NewChatTitleUpdated(chatID, generatedTitle)
	if err := a.notifyApp(ctx, chat.User.ID, json.MustMarshal(ctx, titleUpdated)); err != nil {
		return fmt.Errorf("failed to publish a title updated event: %w", err)
	}

	return nil
}
