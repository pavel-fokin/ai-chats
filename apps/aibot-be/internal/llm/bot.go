package llm

import (
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"

	"pavel-fokin/ai/apps/ai-bot/internal/app"
)

func MessageBot(ctx context.Context, messageText string) (app.Message, error) {
	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return app.Message{}, err
	}

	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, messageText)
	if err != nil {
		return app.Message{}, err
	}

	return app.Message{
		Text: completion,
	}, nil
}
