package domain

import (
	"context"
	"errors"
)

var (
	ErrOllamaModelInvalidStatus = errors.New("invalid status, valid options are: available, pulling")
)

type OllamaModels interface {
	AddModelPullingFinished(ctx context.Context, model string, finalStatus OllamaPullingFinalStatus) error
	AddModelPullingStarted(ctx context.Context, model string) error
	FindOllamaModelsPullingInProgress(ctx context.Context) ([]OllamaModel, error)
}

// OllamaModelStatus is the status of an Ollama model.
type OllamaModelStatus string

const (
	OllamaModelStatusAny       OllamaModelStatus = "any"
	OllamaModelStatusPulling   OllamaModelStatus = "pulling"
	OllamaModelStatusAvailable OllamaModelStatus = "available"
)

// OllamaModelsFilter is the filter for Ollama models.
type OllamaModelsFilter struct {
	Status OllamaModelStatus
}

// NewOllamaModelsFilter creates a new OllamaModelsFilter.
func NewOllamaModelsFilter(status string) (OllamaModelsFilter, error) {
	if status == "" {
		return OllamaModelsFilter{Status: OllamaModelStatusAny}, nil
	}

	modelStatus := OllamaModelStatus(status)
	if !isValidOllamaModelStatus(modelStatus) {
		return OllamaModelsFilter{}, ErrOllamaModelInvalidStatus
	}

	return OllamaModelsFilter{Status: modelStatus}, nil
}

// isValidOllamaModelStatus validates if the provided status is valid.
func isValidOllamaModelStatus(status OllamaModelStatus) bool {
	switch status {
	case OllamaModelStatusPulling, OllamaModelStatusAvailable:
		return true
	default:
		return false
	}
}
