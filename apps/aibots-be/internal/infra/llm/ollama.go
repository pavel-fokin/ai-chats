package llm

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type ChatModel struct {
	llm llms.Model
}

func NewChatModel(model string) (*ChatModel, error) {
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		return nil, err
	}

	return &ChatModel{llm: llm}, nil
}

func (c *ChatModel) GenerateResponse(ctx context.Context, history []domain.Message) (domain.Message, error) {
	content := []llms.MessageContent{}
	for _, message := range history {
		switch message.Sender {
		case "AI":
			content = append(content, llms.TextParts(llms.ChatMessageTypeAI, message.Text))
		case "User":
			content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, message.Text))
		default:
			return domain.Message{}, fmt.Errorf("unknown sender: %s", message.Sender)
		}
	}

	completion, err := c.llm.GenerateContent(ctx, content)
	if err != nil {
		return domain.Message{}, err
	}

	text := completion.Choices[0].Content

	return domain.NewMessage("AI", text), nil
}

func (c *ChatModel) GenerateTitle(ctx context.Context, history []domain.Message) (string, error) {
	content := []llms.MessageContent{}

	for _, message := range history {
		switch message.Sender {
		case "AI":
			content = append(content, llms.TextParts(llms.ChatMessageTypeAI, message.Text))
		case "User":
			content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, message.Text))
		default:
			return "", fmt.Errorf("unknown sender: %s", message.Sender)
		}
	}

	content = append(
		content,
		llms.TextParts(
			llms.ChatMessageTypeHuman,
			"Provide a one-sentence, short title of this conversation. Use less than 100 characters. Don't use quotes or special characters.",
		),
	)

	completion, err := c.llm.GenerateContent(ctx, content)
	if err != nil {
		return "", err
	}

	text := completion.Choices[0].Content

	return text, nil
}
