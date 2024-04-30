package llm

import (
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"

	"pavel-fokin/ai/apps/ai-bot/internal/app"
)

type ChatBot struct {
	llm llms.Model
}

func NewChatModel(model string) (*ChatBot, error) {
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		return nil, err
	}

	return &ChatBot{llm: llm}, nil
}

func (c *ChatBot) SingleMessage(ctx context.Context, prompt string) (app.Message, error) {
	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return app.Message{}, err
	}

	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		return app.Message{}, err
	}

	return app.Message{
		Text: completion,
	}, nil
}

func (c *ChatBot) ChatMessage(ctx context.Context, history []string, prompt string) (app.Message, error) {
	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return app.Message{}, err
	}

	content := []llms.MessageContent{}
	for _, message := range history {
		content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, message))
	}

	content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, prompt))

	completion, err := llm.GenerateContent(ctx, content)
	if err != nil {
		return app.Message{}, err
	}

	text := completion.Choices[0].Content

	return app.Message{
		Text: text,
	}, nil
}
