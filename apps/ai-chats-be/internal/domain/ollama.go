package domain

import "context"

// Models is a repository for Ollama models.
type Ollama interface {
	List(context.Context) ([]Model, error)
	Pull(context.Context, Model) error
	Delete(context.Context, Model) error
}
