package domain

import "context"

// OllamaClient is an interface for the Ollama client.
type OllamaClient interface {
	List(context.Context) ([]OllamaModel, error)
	Pull(context.Context, OllamaModel) error
	Delete(context.Context, OllamaModel) error
}
