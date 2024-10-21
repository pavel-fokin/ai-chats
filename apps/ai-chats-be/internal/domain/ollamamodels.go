package domain

import (
	"context"
	"errors"
)

var (
	ErrOllamaModelInvalidStatus = errors.New("invalid status, valid options are: available, pulling")
)

type OllamaModels interface {
	Save(ctx context.Context, model OllamaModel) error
	FindOllamaModelsPullInProgress(ctx context.Context) ([]OllamaModel, error)
}

// OllamaModelsFilter is the filter for Ollama models.
type OllamaModelsFilter struct {
	Status OllamaModelStatus
}

// NewOllamaModelsFilter creates a new OllamaModelsFilter.
func NewOllamaModelsFilter(status string) (OllamaModelsFilter, error) {
	if status == "" {
		return OllamaModelsFilter{}, nil
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
