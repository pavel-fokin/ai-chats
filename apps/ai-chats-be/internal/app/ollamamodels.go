package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

func (a *App) ListModels(ctx context.Context) ([]domain.Model, error) {
	return a.ollama.List(ctx)
}

func (a *App) PullModel(ctx context.Context, modelName string) error {
	return a.ollama.Pull(ctx, domain.NewModel(modelName, "latest"))
}

func (a *App) DeleteModel(ctx context.Context, modelName string) error {
	return a.ollama.Delete(ctx, domain.NewModel(modelName, "latest"))
}
