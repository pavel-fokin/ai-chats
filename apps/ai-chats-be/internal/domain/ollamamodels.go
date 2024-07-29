package domain

import "context"

type OllamaModels interface {
	AddModelPullingStarted(ctx context.Context, model string) error
	AddModelPullingFinished(ctx context.Context, model string, finalStatus OllamaPullingFinalStatus) error
	AllModelsWithPullingInProgress(ctx context.Context) ([]string, error)
}
