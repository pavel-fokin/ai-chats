package domain

import (
	"context"

	"ai-chats/internal/pkg/types"
)

const (
	OllamaModelPullProgressType types.MessageType = "ollamaModelPullProgress"
)

// OllamaModelPullProgress is a message that represents the progress of a model pulling.
type OllamaModelPullProgress struct {
	Status    string `json:"status"`
	Total     int64  `json:"total"`
	Completed int64  `json:"completed"`
}

func (p OllamaModelPullProgress) Type() types.MessageType {
	return OllamaModelPullProgressType
}

// PullProgressFunc is a function that streams the progress of a model pulling.
type PullProgressFunc func(progress OllamaModelPullProgress) error

// OllamaClient is an interface for the Ollama client.
type OllamaClient interface {
	NewModel(OllamaModel) (Model, error)
	List(context.Context) ([]OllamaModel, error)
	Pull(context.Context, string, PullProgressFunc) error
	Delete(context.Context, string) error
}
