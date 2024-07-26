package app

import (
	"context"
	"errors"
	"fmt"

	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/datatypes"
)

// ListOllamaModels retrieves a list of Ollama models from the Ollama client and performs
// operations to synchronize the models with the added models in the application.
// It returns a slice of domain.OllamaModel and an error if any.
func (a *App) ListOllamaModels(ctx context.Context) ([]domain.OllamaModel, error) {
	// Get all available models from the Ollama client.
	ollamaClientModels, err := a.ollamaClient.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to request ollama models from client: %w", err)
	}

	ollamaClientModelsSet := datatypes.NewSet(
		ollamaClientModels,
		func(ollamaClientModel domain.OllamaClientModel) string {
			return ollamaClientModel.Model
		},
	)

	// Get all added models in the application.
	ollamaAddedModels, err := a.ollamaModels.AllAdded(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all available ollama models: %w", err)
	}

	ollamaAddedModelsSet := datatypes.NewSet(
		ollamaAddedModels,
		func(ollamaModel domain.OllamaModel) string {
			return ollamaModel.Model
		},
	)

	// Delete models that are not available in the Ollama client.
	for _, ollamaAddedModel := range ollamaAddedModels {
		if !ollamaClientModelsSet.Contains(ollamaAddedModel.Model) {
			err := a.ollamaModels.Delete(ctx, ollamaAddedModel)
			if err != nil {
				return nil, fmt.Errorf("failed to delete ollama model: %w", err)
			}
		}
	}

	// Add models that are available in the Ollama client but not added in the application.
	for _, ollamaClientModel := range ollamaClientModels {
		if !ollamaAddedModelsSet.Contains(ollamaClientModel.Model) {
			err := a.ollamaModels.Add(ctx, *domain.NewOllamaModel(ollamaClientModel.Model))
			if err != nil {
				return nil, fmt.Errorf("failed to add ollama model: %w", err)
			}
		}
	}

	// Get final list of added models in the application.
	ollamaModels, err := a.ollamaModels.AllAdded(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all available ollama models: %w", err)
	}

	return ollamaModels, nil
}

func (a *App) PullOllamaModel(ctx context.Context, model string) error {
	ollamaModel, err := a.ollamaModels.Find(ctx, model)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrOllamaModelNotFound):
			ollamaModel = *domain.NewOllamaModel(model)
		default:
			return fmt.Errorf("failed to find ollama model: %w", err)
		}
	}

	ollamaModel.Pull()
	if err := a.ollamaModels.Save(ctx, ollamaModel); err != nil {
		return fmt.Errorf("failed to save ollama model: %w", err)
	}

	if err := a.ollamaClient.Pull(ctx, model); err != nil {
		return fmt.Errorf("failed to pull ollama model: %w", err)
	}

	return nil
}

func (a *App) DeleteOllamaModel(ctx context.Context, model string) error {
	ollamaModel, err := a.ollamaModels.Find(ctx, model)
	if err != nil {
		return fmt.Errorf("failed to find ollama model: %w", err)
	}

	ollamaModel.Delete()
	if err := a.ollamaModels.Delete(ctx, ollamaModel); err != nil {
		return fmt.Errorf("failed to delete ollama model: %w", err)
	}

	return a.ollamaClient.Delete(ctx, model)
}
