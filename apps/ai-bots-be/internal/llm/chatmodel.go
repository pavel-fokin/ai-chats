package llm

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"
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

func (c *ChatModel) SingleMessage(ctx context.Context, prompt string) (domain.Message, error) {
	completion, err := llms.GenerateFromSinglePrompt(ctx, c.llm, prompt)
	if err != nil {
		return domain.Message{}, err
	}

	return domain.Message{
		Actor: domain.Actor{
			Type: domain.AI,
		},
		Text: completion,
	}, nil
}

func (c *ChatModel) ChatMessage(ctx context.Context, history []domain.Message, message string) (domain.Message, error) {
	content := []llms.MessageContent{}
	for _, message := range history {
		switch message.Actor.Type {
		case domain.AI:
			content = append(content, llms.TextParts(llms.ChatMessageTypeAI, message.Text))
		case domain.Human:
			content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, message.Text))
		default:
			return domain.Message{}, fmt.Errorf("unknown actor type: %s", message.Actor.Type)
		}
	}

	content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, message))

	completion, err := c.llm.GenerateContent(ctx, content)
	if err != nil {
		return domain.Message{}, err
	}

	if len(completion.Choices) == 0 {
		return domain.Message{}, fmt.Errorf("no completion choices")
	}
	text := completion.Choices[0].Content

	return domain.Message{
		Actor: domain.Actor{
			Type: domain.AI,
		},
		Text: text,
	}, nil
}
