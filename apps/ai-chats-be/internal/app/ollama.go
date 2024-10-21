package app

import (
	"context"
	"fmt"

	"ai-chats/internal/domain"
)

// FindOllamaModels retrieves Ollama models based on the provided filter.
func (a *App) FindOllamaModels(ctx context.Context, filter domain.OllamaModelsFilter) ([]domain.OllamaModel, error) {
	if filter.Status == domain.OllamaModelStatusPulling {
		return a.findPullingOllamaModels(ctx)
	}

	if filter.Status == domain.OllamaModelStatusAvailable {
		return a.findAvailableOllamaModels(ctx)
	}

	// Find ollama models with any status.
	var ollamaModels []domain.OllamaModel

	pullingOllamaModels, err := a.findPullingOllamaModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ollama models with pulling in progress: %w", err)
	}

	availableOllamaModels, err := a.findAvailableOllamaModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ollama models available: %w", err)
	}

	ollamaModels = append(ollamaModels, pullingOllamaModels...)
	ollamaModels = append(ollamaModels, availableOllamaModels...)

	return ollamaModels, nil
}

// findDescription finds the description of an Ollama model.
func (a *App) findDescription(ctx context.Context, ollamaModel domain.OllamaModel) (domain.OllamaModel, error) {
	description, err := a.modelsLibrary.FindDescription(ctx, ollamaModel.Name)
	if err != nil {
		description = "Description is not available."
	}

	ollamaModel.Description = description

	return ollamaModel, nil
}

func (a *App) findPullingOllamaModels(ctx context.Context) ([]domain.OllamaModel, error) {
	var ollamaModels []domain.OllamaModel

	pullingOllamaModels, err := a.ollamaModels.FindOllamaModelsPullInProgress(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ollama models with pulling in progress: %w", err)
	}

	for _, ollamaModel := range pullingOllamaModels {
		fmt.Println("ollamaModel", ollamaModel)
		ollamaModel, err := a.findDescription(ctx, ollamaModel)
		if err != nil {
			return nil, fmt.Errorf("failed to create ollama model: %w", err)
		}

		ollamaModels = append(ollamaModels, ollamaModel)
	}

	return ollamaModels, nil
}

func (a *App) findAvailableOllamaModels(ctx context.Context) ([]domain.OllamaModel, error) {
	var ollamaModels []domain.OllamaModel

	availableOllamaModels, err := a.ollamaClient.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to request ollama models from client: %w", err)
	}

	for _, ollamaModel := range availableOllamaModels {
		ollamaModel, err := a.findDescription(ctx, ollamaModel)
		if err != nil {
			return nil, fmt.Errorf("failed to create ollama model: %w", err)
		}

		ollamaModels = append(ollamaModels, ollamaModel)
	}

	return ollamaModels, nil
}

func (a *App) PullOllamaModelAsync(ctx context.Context, model string) error {
	pullOllamaModelCommand := NewPullOllamaModel(model)
	if err := a.pubsub.Publish(
		ctx,
		PullOllamaModelTopic,
		pullOllamaModelCommand,
	); err != nil {
		return fmt.Errorf("failed to publish pull ollama model command: %w", err)
	}

	return nil
}

// PullOllamaModelJob pulls an Ollama model.
func (a *App) PullOllamaModel(ctx context.Context, model string) error {
	ollamaModel, _ := domain.NewOllamaModel(model)
	ollamaModel.PullStarted()
	if err := a.ollamaModels.Save(ctx, ollamaModel); err != nil {
		return fmt.Errorf("failed to add ollama model pulling started: %w", err)
	}
	ollamaModel.ClearEvents()

	progressFunc := func(progress OllamaModelPullProgress) error {
		if err := a.notifyOllamaModelPulling(ctx, model, progress); err != nil {
			return fmt.Errorf("failed to notify ollama model pulling progress: %w", err)
		}
		return nil
	}

	if err := a.ollamaClient.Pull(ctx, model, progressFunc); err != nil {
		ollamaModel.PullFailed()
		if err := a.ollamaModels.Save(ctx, ollamaModel); err != nil {
			return fmt.Errorf("failed to add ollama model pulling failed: %w", err)
		}
		ollamaModel.ClearEvents()

		return fmt.Errorf("failed to pull ollama model: %w", err)
	}

	ollamaModel.PullCompleted()
	if err := a.ollamaModels.Save(ctx, ollamaModel); err != nil {
		return fmt.Errorf("failed to add ollama model pulling completed: %w", err)
	}
	ollamaModel.ClearEvents()

	return nil
}

// DeleteOllamaModel deletes an Ollama model.
func (a *App) DeleteOllamaModel(ctx context.Context, model string) error {
	return a.ollamaClient.Delete(ctx, model)
}
