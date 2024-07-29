package app

import (
	"context"
	"fmt"

	"ai-chats/internal/app/commands"
	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/json"
	"ai-chats/internal/worker"
)

// ListOllamaModels retrieves a list of Ollama models from the Ollama client.
func (a *App) ListOllamaModels(ctx context.Context) ([]domain.OllamaModel, error) {
	var ollamaModels []domain.OllamaModel

	ollamaModelsStrings, err := a.ollamaModels.AllModelsWithPullingInProgress(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ollama models with pulling in progress: %w", err)
	}

	for _, model := range ollamaModelsStrings {
		description, err := a.models.FindDescription(ctx, model)
		if err != nil {
			description = "Description is not available."
		}

		ollamaModel := domain.NewOllamaModel(model, description)
		ollamaModel.IsPulling = true
		ollamaModels = append(ollamaModels, ollamaModel)
	}

	ollamaClientModels, err := a.ollamaClient.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to request ollama models from client: %w", err)
	}

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

func (a *App) PullOllamaModelAsync(ctx context.Context, model string) error {
	pullOllamaModelCommand := commands.NewPullOllamaModel(model)
	if err := a.pubsub.Publish(
		ctx,
		worker.PullOllamaModelTopic,
		json.MustMarshal(ctx, pullOllamaModelCommand),
	); err != nil {
		return fmt.Errorf("failed to publish pull ollama model command: %w", err)
	}

	return nil
}

func (a *App) DeleteOllamaModel(ctx context.Context, model string) error {
	return a.ollamaClient.Delete(ctx, model)
}
