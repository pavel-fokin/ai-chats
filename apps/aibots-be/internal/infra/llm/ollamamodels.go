package llm

import (
	"context"
	"fmt"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"strings"

	"github.com/ollama/ollama/api"
)

type OllamaModels struct {
	client *api.Client
}

func NewOllamaModels() (*OllamaModels, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("failed to create a client: %w", err)
	}

	return &OllamaModels{client: client}, nil
}

func (o *OllamaModels) All(ctx context.Context) ([]domain.Model, error) {
	resp, err := o.client.List(ctx)
	if err != nil {
		return nil, err
	}

	models := []domain.Model{}
	for _, model := range resp.Models {
		name := strings.Split(model.Name, ":")[0]
		tag := strings.Split(model.Name, ":")[1]
		models = append(models, domain.NewModel(name, tag))
	}

	return models, nil
}

func (o *OllamaModels) Pull(ctx context.Context, model domain.Model) error {
	req := &api.PullRequest{
		Name: model.Name + ":" + model.Tag,
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

func (o *OllamaModels) Delete(ctx context.Context, model domain.Model) error {
	req := &api.DeleteRequest{
		Name: model.Name + ":" + model.Tag,
	}

	if err := o.client.Delete(ctx, req); err != nil {
		return err
	}

	return nil
}
