package domain

import "context"

// Models is a repository for Ollama models.
type Ollama interface {
	List(context.Context) ([]OllamaModel, error)
	Pull(context.Context, OllamaModel) error
	Delete(context.Context, OllamaModel) error
}
