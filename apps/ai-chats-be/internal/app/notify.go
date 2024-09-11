package app

import (
	"context"
	"fmt"
)

func (a *App) notifyInChat(ctx context.Context, chatID string, event []byte) error {
	if err := a.pubsub.Publish(ctx, chatID, event); err != nil {
		return fmt.Errorf("failed to publish a message sent event: %w", err)
	}
	return nil
}
