package ollama

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ollama/ollama/api"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
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

func (o *OllamaModels) List(ctx context.Context) ([]domain.Model, error) {
	resp, err := o.client.List(ctx)
	if err != nil {
		return nil, err
	}

	models := []domain.Model{}
	for _, model := range resp.Models {
		fmt.Println("model.Name", model.Name)
		fmt.Println("model.Model", model.Model)
		fmt.Println("model.ModifiedAt", model.ModifiedAt)
		fmt.Println("model.Size", model.Size)
		fmt.Println("model.Digest", model.Digest)
		fmt.Println("model.Details", model.Details)
		fmt.Println("")
		name := strings.Split(model.Model, ":")[0]
		tag := strings.Split(model.Model, ":")[1]
		models = append(models, domain.NewModel(name, tag))
	}

	return models, nil
}

func (o *OllamaModels) Pull(ctx context.Context, model domain.Model) error {
	req := &api.PullRequest{
		Model: model.Name + ":" + model.Tag,
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
		Model: model.Name + ":" + model.Tag,
	}

	if err := o.client.Delete(ctx, req); err != nil {
		return err
	}

	return nil
}
