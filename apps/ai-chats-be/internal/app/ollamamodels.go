package app

import (
	"context"

	"ai-chats/internal/domain"
)

func (a *App) ListModels(ctx context.Context) ([]domain.OllamaModel, error) {
	return a.ollama.List(ctx)
}

func (a *App) PullModel(ctx context.Context, model string) error {
	return a.ollama.Pull(ctx, domain.NewOllamaModel(model))
}

func (a *App) DeleteModel(ctx context.Context, model string) error {
	return a.ollama.Delete(ctx, domain.NewOllamaModel(model))
}
