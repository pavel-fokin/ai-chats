package app

import (
	"context"

	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

const (
	OllamaModelPullingProgressType types.MessageType = "ollamaModelPullingProgress"
)

// OllamaModelPullingProgress is a message that represents the progress of a model pulling.
type OllamaModelPullingProgress struct {
	Status    string `json:"status"`
	Total     int64  `json:"total"`
	Completed int64  `json:"completed"`
}

func (p OllamaModelPullingProgress) Type() types.MessageType {
	return OllamaModelPullingProgressType
}

// PullingFunc is a function that streams the progress of a model pulling.
type PullingFunc func(progress OllamaModelPullingProgress) error

// OllamaClient is an interface for the Ollama client.
type OllamaClient interface {
	NewModel(domain.OllamaModel) (domain.Model, error)
	List(context.Context) ([]domain.OllamaModel, error)
	Pull(context.Context, string, PullingFunc) error
	Delete(context.Context, string) error
}
