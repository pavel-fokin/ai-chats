package app

import (
	"context"
	"fmt"

	"ai-chats/internal/domain"
	"ai-chats/internal/domain/events"
	"ai-chats/internal/pkg/json"
)

// ProcessAddedMessage processes a message added event.
func (a *App) ProcessAddedMessage(ctx context.Context, event events.MessageAdded) error {

	if err := a.pubsub.Publish(ctx, event.ChatID.String(), json.MustMarshal(ctx, event)); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}

	switch {
	case event.Message.IsFromUser():
		a.GenerateResponse(ctx, event.ChatID)
	case event.Message.IsFromModel():
		// Ignore messages from models.
	default:
		return fmt.Errorf("unknown message type: %s", event.Message)
	}

	messages, err := a.chats.AllMessages(ctx, event.ChatID)
	if err != nil {
		return fmt.Errorf("failed to find a chat: %w", err)
	}

	if len(messages) == 2 {
		return a.GenerateTitle(ctx, event.ChatID)
	}

	return nil
}

// PullOllamaModelJob pulls an Ollama model.
func (a *App) PullOllamaModelJob(ctx context.Context, model string) error {
	ollamaModel := domain.NewOllamaModel(model, "")

	err := a.ollamaModels.AddModelPullingStarted(ctx, ollamaModel.Name())
	if err != nil {
		return fmt.Errorf("failed to add ollama model pulling started: %w", err)
	}

	if err := a.ollamaClient.Pull(ctx, model); err != nil {
		if err := a.ollamaModels.AddModelPullingFinished(
			ctx,
			ollamaModel.Name(),
			domain.OllamaPullingFinalStatusFailed,
		); err != nil {
			return fmt.Errorf("failed to add ollama model pulling finished: %w", err)
		}

		return fmt.Errorf("failed to pull ollama model: %w", err)
	}

	if err := a.ollamaModels.AddModelPullingFinished(
		ctx,
		ollamaModel.Name(),
		domain.OllamaPullingFinalStatusSuccess,
	); err != nil {
		return fmt.Errorf("failed to add ollama model pulling finished: %w", err)
	}

	return nil
}
