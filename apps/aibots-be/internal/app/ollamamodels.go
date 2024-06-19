package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

func (a *App) ListModels(ctx context.Context) ([]domain.Model, error) {
	return a.ollama.List(ctx)
}
