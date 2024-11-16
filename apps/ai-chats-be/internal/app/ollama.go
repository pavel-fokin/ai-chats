package app

import (
	"context"
	"fmt"

	"ai-chats/internal/domain"
)

// Ollama is an Ollama service.
type Ollama struct {
	ollamaClient  domain.OllamaClient
	ollamaModels  domain.OllamaModels
	modelsLibrary domain.ModelsLibrary
	pubsub        PubSub
}

// NewOllama creates a new Ollama service.
func NewOllama(
	ollamaClient domain.OllamaClient,
	ollamaModels domain.OllamaModels,
	modelsLibrary domain.ModelsLibrary,
	pubsub PubSub,
) *Ollama {
	return &Ollama{ollamaClient: ollamaClient, ollamaModels: ollamaModels, modelsLibrary: modelsLibrary, pubsub: pubsub}
}

// GetOllamaModelsLibrary retrieves the library of Ollama models.
func (o *Ollama) GetOllamaModelsLibrary(ctx context.Context) ([]*domain.ModelCard, error) {
	return o.modelsLibrary.FindAll(ctx)
}

// FindOllamaModels retrieves Ollama models based on the provided filter.
func (o *Ollama) FindOllamaModels(ctx context.Context, filter domain.OllamaModelsFilter) ([]domain.OllamaModel, error) {
	if filter.Status == domain.OllamaModelStatusPulling {
		return o.findPullingOllamaModels(ctx)
	}

	if filter.Status == domain.OllamaModelStatusAvailable {
		return o.findAvailableOllamaModels(ctx)
	}

	// Find ollama models with any status.
	var ollamaModels []domain.OllamaModel

	pullingOllamaModels, err := o.findPullingOllamaModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ollama models with pulling in progress: %w", err)
	}

	availableOllamaModels, err := o.findAvailableOllamaModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ollama models available: %w", err)
	}

	ollamaModels = append(ollamaModels, pullingOllamaModels...)
	ollamaModels = append(ollamaModels, availableOllamaModels...)

	return ollamaModels, nil
}

// findDescription finds the description of an Ollama model.
func (o *Ollama) findDescription(ctx context.Context, ollamaModel domain.OllamaModel) (domain.OllamaModel, error) {
	description, err := o.modelsLibrary.FindDescription(ctx, ollamaModel.Name)
	if err != nil {
		description = "Description is not available."
	}

	ollamaModel.Description = description

	return ollamaModel, nil
}

func (o *Ollama) findPullingOllamaModels(ctx context.Context) ([]domain.OllamaModel, error) {
	var ollamaModels []domain.OllamaModel

	pullingOllamaModels, err := o.ollamaModels.FindOllamaModelsPullInProgress(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ollama models with pulling in progress: %w", err)
	}

	for _, ollamaModel := range pullingOllamaModels {
		fmt.Println("ollamaModel", ollamaModel)
		ollamaModel, err := o.findDescription(ctx, ollamaModel)
		if err != nil {
			return nil, fmt.Errorf("failed to create ollama model: %w", err)
		}

		ollamaModels = append(ollamaModels, ollamaModel)
	}

	return ollamaModels, nil
}

func (o *Ollama) findAvailableOllamaModels(ctx context.Context) ([]domain.OllamaModel, error) {
	var ollamaModels []domain.OllamaModel

	availableOllamaModels, err := o.ollamaClient.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to request ollama models from client: %w", err)
	}

	for _, ollamaModel := range availableOllamaModels {
		ollamaModel, err := o.findDescription(ctx, ollamaModel)
		if err != nil {
			return nil, fmt.Errorf("failed to create ollama model: %w", err)
		}

		ollamaModels = append(ollamaModels, ollamaModel)
	}

	return ollamaModels, nil
}

func (o *Ollama) PullOllamaModelAsync(ctx context.Context, model string) error {
	pullOllamaModelCommand := NewPullOllamaModel(model)
	if err := o.pubsub.Publish(
		ctx,
		PullOllamaModelTopic,
		pullOllamaModelCommand,
	); err != nil {
		return fmt.Errorf("failed to publish pull ollama model command: %w", err)
	}

	return nil
}

// PullOllamaModelJob pulls an Ollama model.
func (o *Ollama) PullOllamaModel(ctx context.Context, model string) error {
	ollamaModel, _ := domain.NewOllamaModel(model)
	ollamaModel.PullStarted()
	if err := o.ollamaModels.Save(ctx, ollamaModel); err != nil {
		return fmt.Errorf("failed to add ollama model pulling started: %w", err)
	}
	ollamaModel.ClearEvents()

	progressFunc := func(progress domain.OllamaModelPullProgress) error {
		if err := o.notifyOllamaModelPulling(ctx, model, progress); err != nil {
			return fmt.Errorf("failed to notify ollama model pulling progress: %w", err)
		}
		return nil
	}

	if err := o.ollamaClient.Pull(ctx, model, progressFunc); err != nil {
		ollamaModel.PullFailed()
		if err := o.ollamaModels.Save(ctx, ollamaModel); err != nil {
			return fmt.Errorf("failed to add ollama model pulling failed: %w", err)
		}
		ollamaModel.ClearEvents()

		return fmt.Errorf("failed to pull ollama model %s: %w", model, err)
	}

	ollamaModel.PullCompleted()
	if err := o.ollamaModels.Save(ctx, ollamaModel); err != nil {
		return fmt.Errorf("failed to add ollama model pulling completed: %w", err)
	}
	ollamaModel.ClearEvents()

	return nil
}

// DeleteOllamaModel deletes an Ollama model.
func (o *Ollama) DeleteOllamaModel(ctx context.Context, model string) error {
	return o.ollamaClient.Delete(ctx, model)
}
