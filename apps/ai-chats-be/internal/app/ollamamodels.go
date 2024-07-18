package app

import (
	"context"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

func (a *App) ListModels(ctx context.Context) ([]domain.OllamaModel, error) {
	return a.ollama.List(ctx)
}

func (a *App) PullModel(ctx context.Context, model string) error {
	return a.ollama.Pull(ctx, domain.NewModel(model))
}

func (a *App) DeleteModel(ctx context.Context, model string) error {
	return a.ollama.Delete(ctx, domain.NewModel(model))
}
