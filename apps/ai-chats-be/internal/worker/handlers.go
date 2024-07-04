package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain/events"
)

const (
	MessageAddedTopic Topic = "message-added"
)

func (w *Worker) SetupHandlers(app App) {
	w.RegisterHandler(MessageAddedTopic, 1, w.MessageAdded(app))
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
