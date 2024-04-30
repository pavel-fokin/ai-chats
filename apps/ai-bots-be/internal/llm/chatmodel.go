package llm

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app"
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

func (c *ChatModel) SingleMessage(ctx context.Context, prompt string) (app.Message, error) {
	completion, err := llms.GenerateFromSinglePrompt(ctx, c.llm, prompt)
	if err != nil {
		return app.Message{}, err
	}

	return app.Message{
		Text: completion,
	}, nil
}

func (c *ChatModel) ChatMessage(ctx context.Context, history []domain.Message, message string) (app.Message, error) {
	content := []llms.MessageContent{}
	for _, message := range history {
		switch message.Actor.Type {
		case "ai":
			content = append(content, llms.TextParts(llms.ChatMessageTypeAI, message.Text))
		case "user":
			content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, message.Text))
		default:
			return app.Message{}, fmt.Errorf("unknown actor type: %s", message.Actor.Type)
		}
	}

	content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, message))

	completion, err := c.llm.GenerateContent(ctx, content)
	if err != nil {
		return app.Message{}, err
	}

	if len(completion.Choices) == 0 {
		return app.Message{}, fmt.Errorf("no completion choices")
	}
	text := completion.Choices[0].Content

	return app.Message{
		Text: text,
	}, nil
}
