package domain

import (
	"context"

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

// PullingStreamFunc is a function that streams the progress of a model pulling.
type PullingStreamFunc func(progress OllamaModelPullingProgress) error

// OllamaClient is an interface for the Ollama client.
type OllamaClient interface {
	List(context.Context) ([]OllamaModel, error)
	Pull(context.Context, string, PullingStreamFunc) error
	Delete(context.Context, string) error
}
