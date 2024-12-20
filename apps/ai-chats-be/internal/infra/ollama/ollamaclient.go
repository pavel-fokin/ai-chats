package ollama

import (
	"context"
	"log"

	"github.com/ollama/ollama/api"

	"ai-chats/internal/domain"
)

type OllamaClient struct {
	client *api.Client
}

func NewOllamaClient() *OllamaClient {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatalf("Failed to create Ollama client: %v", err)
	}

	return &OllamaClient{client: client}
}

// NewModel creates a new model to interact with the given model.
func (o *OllamaClient) NewModel(model domain.OllamaModel) (domain.Model, error) {
	return NewModel(o.client, model)
}

// List returns the list of available models.
func (o *OllamaClient) List(ctx context.Context) ([]domain.OllamaModel, error) {
	resp, err := o.client.List(ctx)
	if err != nil {
		return nil, err
	}

	models := []domain.OllamaModel{}
	for _, model := range resp.Models {
		ollamaModel, _ := domain.NewOllamaModel(model.Model)
		ollamaModel.SetStatus(domain.OllamaModelStatusAvailable)
		models = append(models, ollamaModel)
	}

	return models, nil
}

// Pull sends request to the Ollama server to pull a model and streams the progress to the given function.
func (o *OllamaClient) Pull(ctx context.Context, model string, fn domain.PullProgressFunc) error {
	req := &api.PullRequest{
		Model: model,
	}

	progressFunc := func(resp api.ProgressResponse) error {
		progress := domain.OllamaModelPullProgress{
			Status:    resp.Status,
			Total:     resp.Total,
			Completed: resp.Completed,
		}

		return fn(progress)
	}

	return o.client.Pull(ctx, req, progressFunc)
}

// Delete deletes the given model.
func (o *OllamaClient) Delete(ctx context.Context, model string) error {
	req := &api.DeleteRequest{
		Model: model,
	}

	if err := o.client.Delete(ctx, req); err != nil {
		return err
	}

	return nil
}
