package app

import (
	"context"

	"ai-chats/internal/domain"
)

func (a *App) AllModelCards(ctx context.Context) ([]domain.ModelCard, error) {
	return a.models.AllModelCards(ctx)
}
