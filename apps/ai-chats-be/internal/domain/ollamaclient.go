package domain

import (
	"context"
	"strings"
)

// OllamaClient is an interface for the Ollama client.
type OllamaClient interface {
	List(context.Context) ([]OllamaClientModel, error)
	Pull(context.Context, string) error
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
