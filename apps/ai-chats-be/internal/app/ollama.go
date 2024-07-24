package app

import (
	"context"
	"fmt"

	"ai-chats/internal/domain"
)

func (a *App) ListModels(ctx context.Context) ([]domain.OllamaModel, error) {

	ollamaClientModels, err := a.ollamaClient.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list ollame models: %w", err)
	}

	var ollamaModels []domain.OllamaModel
	for _, clientModel := range ollamaClientModels {
		modelCard, err := a.models.FindModelCard(ctx, clientModel.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to find model cards: %w", err)
		}
		ollamaModels = append(ollamaModels, domain.OllamaModel{
			Model:       clientModel.Model,
			Description: modelCard.Description,
		})
	}

	return ollamaModels, nil
}

func (a *App) PullModel(ctx context.Context, model string) error {
	return a.ollamaClient.Pull(ctx, model)
}

func (a *App) DeleteModel(ctx context.Context, model string) error {
	return a.ollamaClient.Delete(ctx, model)
}
