package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

func (a *App) AllModelCards(ctx context.Context) ([]domain.ModelCard, error) {
	return a.models.AllModelCards(ctx)
}
