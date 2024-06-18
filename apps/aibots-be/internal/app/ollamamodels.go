package app

import (
	"context"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

func (a *App) AllOllamaModels(ctx context.Context) ([]domain.Model, error) {
	return a.models.All(ctx)
}
