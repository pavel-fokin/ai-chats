package ollama

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"

	"ai-chats/internal/domain"
)

type Model struct {
	client *api.Client
	model  domain.OllamaModel
}

func NewModel(client *api.Client, model domain.OllamaModel) (*Model, error) {
	return &Model{client: client, model: model}, nil
}

func (m *Model) Chat(ctx context.Context, messages []domain.Message, fn domain.ChatResponseFunc) (domain.Message, error) {
	apiMessages := []api.Message{}
	for _, message := range messages {
		role := ""
		switch {
		case message.IsFromModel():
			role = "assistant"
		case message.IsFromUser():
			role = "user"
		case message.IsFromSystem():
			role = "system"
		default:
			return domain.Message{}, fmt.Errorf("unknown sender: %s", message.Sender)
		}
		apiMessages = append(apiMessages, api.Message{
			Role:    role,
			Content: message.Text,
		})
	}

	req := &api.ChatRequest{
		Model:    m.model.String(),
		Messages: apiMessages,
	}

	modelID := domain.NewModelID(m.model.String())
	message := domain.NewModelMessage(modelID, "")
	respFunc := func(resp api.ChatResponse) error {
		message.Text += resp.Message.Content
		if err := fn(message); err != nil {
			return err
		}

		return nil
	}

	if err := m.client.Chat(ctx, req, respFunc); err != nil {
		return domain.Message{}, err
	}

	return message, nil
}
