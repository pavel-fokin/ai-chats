package app

import (
	"context"
	"fmt"

	"ai-chats/internal/pkg/types"
)

// notifyOllamaModelPulling notifies the ollama model about an event.
func (o *Ollama) notifyOllamaModelPulling(ctx context.Context, model string, event types.Message) error {
	if err := o.pubsub.Publish(ctx, model, event); err != nil {
		return fmt.Errorf("failed to publish an event: %w", err)
	}
	return nil
}
