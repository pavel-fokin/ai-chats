package app

import (
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func SendMessage(ctx context.Context, messageText string) (Message, error) {
	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return Message{}, err
	}

	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, messageText)
	if err != nil {
		return Message{}, err
	}

	return Message{
		Text: completion,
	}, nil
}
