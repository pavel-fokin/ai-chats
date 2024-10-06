package domain

import "context"

type OllamaModels interface {
	AddModelPullingFinished(ctx context.Context, model string, finalStatus OllamaPullingFinalStatus) error
	AddModelPullingStarted(ctx context.Context, model string) error
	FindOllamaModelsPullingInProgress(ctx context.Context) ([]OllamaModel, error)
}
