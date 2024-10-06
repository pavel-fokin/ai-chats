package domain

import (
	"context"
	"strings"

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
	List(context.Context) ([]OllamaClientModel, error)
	Pull(context.Context, string, PullingStreamFunc) error
	Delete(context.Context, string) error
}

type OllamaClientModel struct {
	Model string
}

// NewOllamaClientModel creates a new Ollama client model.
func NewOllamaClientModel(model string) OllamaClientModel {
	return OllamaClientModel{Model: model}
}

func (om OllamaClientModel) Name() string {
	parts := strings.Split(om.Model, ":")
	if len(parts) == 2 {
		return parts[0]
	}
	return om.Model
}
