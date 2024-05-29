package llm

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain/events"
)

type LLM struct {
	client *api.Client
	model  string
}

type StreamFunc func(messageChunk events.MessageChunkReceived) error

func NewOllama(model string) (*LLM, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("failed to create a client: %w", err)
	}

	return &LLM{client: client, model: model}, nil
}

func (l *LLM) GenerateResponse(ctx context.Context, history []domain.Message) (domain.Message, error) {
	messages := []api.Message{}
	for _, message := range history {
		role := ""
		switch message.Sender {
		case "AI":
			role = "assistant"
		case "User":
			role = "user"
		default:
			return domain.Message{}, fmt.Errorf("unknown sender: %s", message.Sender)
		}
		messages = append(messages, api.Message{
			Role:    role,
			Content: message.Text,
		})
	}

	req := &api.ChatRequest{
		Model:    l.model,
		Messages: messages,
		Stream:   new(bool),
	}

	llmMessage := domain.Message{}
	respFunc := func(resp api.ChatResponse) error {
		llmMessage = domain.NewMessage("AI", resp.Message.Content)
		return nil
	}

	if err := l.client.Chat(ctx, req, respFunc); err != nil {
		return domain.Message{}, err
	}

	return llmMessage, nil
}

func (l *LLM) GenerateResponseWithStream(
	ctx context.Context,
	history []domain.Message,
	fn StreamFunc,
) (domain.Message, error) {
	messages := []api.Message{}
	for _, message := range history {
		role := ""
		switch message.Sender {
		case "AI":
			role = "assistant"
		case "User":
			role = "user"
		default:
			return domain.Message{}, fmt.Errorf("unknown sender: %s", message.Sender)
		}
		messages = append(messages, api.Message{
			Role:    role,
			Content: message.Text,
		})
	}

	req := &api.ChatRequest{
		Model:    l.model,
		Messages: messages,
	}

	llmMessage := domain.NewMessage("AI", "")
	respFunc := func(resp api.ChatResponse) error {
		llmMessage.Text += resp.Message.Content

		messageChunkReceived := events.NewMessageChunkReceived(
			llmMessage.ID,
			llmMessage.Sender,
			llmMessage.Text,
			resp.Done,
		)
		if err := fn(messageChunkReceived); err != nil {
			return err
		}

		return nil
	}

	if err := l.client.Chat(ctx, req, respFunc); err != nil {
		return domain.Message{}, err
	}

	return llmMessage, nil
}

func (l *LLM) GenerateTitle(ctx context.Context, history []domain.Message) (string, error) {
	messages := []api.Message{}
	for _, message := range history {
		role := ""
		switch message.Sender {
		case "AI":
			role = "assistant"
		case "User":
			role = "user"
		default:
			return "", fmt.Errorf("unknown sender: %s", message.Sender)
		}
		messages = append(messages, api.Message{
			Role:    role,
			Content: message.Text,
		})
	}

	messages = append(messages, api.Message{
		Role:    "user",
		Content: "Provide a one-sentence, short title of this following conversation. Use less than 100 characters. Don't use quotes or special characters.",
	})

	req := &api.ChatRequest{
		Model:    l.model,
		Messages: messages,
		Stream:   new(bool),
	}

	var title string
	respFunc := func(resp api.ChatResponse) error {
		title = resp.Message.Content
		return nil
	}

	if err := l.client.Chat(ctx, req, respFunc); err != nil {
		return "", err
	}

	return title, nil
}
