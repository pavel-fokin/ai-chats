package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"ai-chats/internal/app/commands"
	"ai-chats/internal/domain/events"
)

const (
	MessageAddedTopic    Topic = "message-added"
	PullOllamaModelTopic Topic = "pull-ollama-model"
)

func (w *Worker) SetupHandlers(app App) {
	w.RegisterHandler(MessageAddedTopic, 1, w.MessageAdded(app))
	w.RegisterHandler(PullOllamaModelTopic, 1, w.PullOllamaModel(app))
}

func (w *Worker) MessageAdded(app App) HandlerFunc {
	return func(ctx context.Context, e []byte) error {
		var messageAdded events.MessageAdded
		if err := json.Unmarshal(e, &messageAdded); err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}

		err := app.ProcessAddedMessage(ctx, messageAdded)
		if err != nil {
			return fmt.Errorf("failed to handle a message added event: %w", err)
		}

		return nil
	}
}

func (w *Worker) PullOllamaModel(app App) HandlerFunc {
	return func(ctx context.Context, e []byte) error {
		var pullOllamaModel commands.PullOllamaModel
		if err := json.Unmarshal(e, &pullOllamaModel); err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}

		err := app.PullOllamaModelJob(ctx, pullOllamaModel.Model)
		if err != nil {
			return fmt.Errorf("failed to pull ollama model: %w", err)
		}

		return nil
	}
}
