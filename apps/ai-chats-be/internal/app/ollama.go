package app

import (
	"context"
	"fmt"

	"ai-chats/internal/domain"
)

// ListOllamaModels retrieves a list of Ollama models from the Ollama client.
func (a *App) ListOllamaModels(ctx context.Context) ([]domain.OllamaModel, error) {
	ollamaClientModels, err := a.ollamaClient.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to request ollama models from client: %w", err)
	}

	var ollamaModels []domain.OllamaModel
	for _, ollamaClientModel := range ollamaClientModels {
		description, err := a.models.FindDescription(ctx, ollamaClientModel.Name())
		if err != nil {
			description = "Description is not available."
		}

		ollamaModel := domain.NewOllamaModel(ollamaClientModel.Model, description)
		ollamaModels = append(ollamaModels, ollamaModel)
	}

	return ollamaModels, nil
}

func (a *App) PullOllamaModel(ctx context.Context, model string) error {
	if err := a.ollamaClient.Pull(ctx, model); err != nil {
		return fmt.Errorf("failed to pull ollama model: %w", err)
	}

	return nil
}

func (a *App) DeleteOllamaModel(ctx context.Context, model string) error {
	return a.ollamaClient.Delete(ctx, model)
}
