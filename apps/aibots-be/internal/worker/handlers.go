package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

func (w *Worker) SetupHandlers(app App) {
	w.RegisterHandler("worker", 1, w.MessageSent(app))
}

func (w *Worker) MessageSent(app App) HandlerFunc {
	return func(ctx context.Context, e []byte) error {
		var messageSent domain.MessageSent
		if err := json.Unmarshal(e, &messageSent); err != nil {
			slog.ErrorContext(w.ctx, "failed to unmarshal event", "err", err)
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}

		err := app.GenerateResponse(w.ctx, messageSent.ChatID)
		if err != nil {
			slog.ErrorContext(w.ctx, "failed to generate a response", "err", err)
			return fmt.Errorf("failed to generate a response: %w", err)
		}

		return nil
	}
}
