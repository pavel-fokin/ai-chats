package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/commands"
)

func (w *Worker) SetupHandlers(app App) {
	w.RegisterHandler("worker", 1, w.GenerateResponse(app))
	w.RegisterHandler("generate-title", 1, w.GenerateTitle(app))
}

func (w *Worker) GenerateResponse(app App) HandlerFunc {
	return func(ctx context.Context, e []byte) error {
		var generateResponse commands.GenerateResponse
		if err := json.Unmarshal(e, &generateResponse); err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}

		err := app.GenerateResponse(ctx, generateResponse.ChatID)
		if err != nil {
			return fmt.Errorf("failed to generate a response: %w", err)
		}

		return nil
	}
}

func (w *Worker) GenerateTitle(app App) HandlerFunc {
	return func(ctx context.Context, e []byte) error {
		var generateTitle commands.GenerateTitle
		if err := json.Unmarshal(e, &generateTitle); err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}

		err := app.GenerateTitle(ctx, generateTitle.ChatID)
		if err != nil {
			return fmt.Errorf("failed to generate a title: %w", err)
		}

		return nil
	}
}
