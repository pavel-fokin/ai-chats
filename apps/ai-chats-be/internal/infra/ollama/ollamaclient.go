package ollama

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"

	"ai-chats/internal/domain"
)

type OllamaModels struct {
	client *api.Client
}

func NewOllamaModels() *OllamaModels {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatalf("Failed to create Ollama client: %v", err)
	}

	return &OllamaModels{client: client}
}

func (o *OllamaModels) List(ctx context.Context) ([]domain.OllamaClientModel, error) {
	resp, err := o.client.List(ctx)
	if err != nil {
		return nil, err
	}

	models := []domain.OllamaClientModel{}
	for _, model := range resp.Models {
		models = append(models, domain.NewOllamaClientModel(model.Model))
	}

	return models, nil
}

func (o *OllamaModels) Pull(ctx context.Context, model string) error {
	req := &api.PullRequest{
		Model: model,
	}

	progressFunc := func(resp api.ProgressResponse) error {
		fmt.Println(resp)
		return nil
	}

	if err := o.client.Pull(ctx, req, progressFunc); err != nil {
		return err
	}

	return nil
}

func (o *OllamaModels) Delete(ctx context.Context, model string) error {
	req := &api.DeleteRequest{
		Model: model,
	}

	if err := o.client.Delete(ctx, req); err != nil {
		return err
	}

	return nil
}
