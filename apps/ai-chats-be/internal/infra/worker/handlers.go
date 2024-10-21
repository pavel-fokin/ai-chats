package worker

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"ai-chats/internal/app"
	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

func (w *Worker) SetupHandlers(a App) {
	w.RegisterHandler(app.GenerateChatTitleTopic, 1, w.GenerateChatTitle(a))
	w.RegisterHandler(app.MessageAddedTopic, 1, w.MessageAdded(a))
	w.RegisterHandler(app.PullOllamaModelTopic, 1, w.PullOllamaModel(a))
}

func (w *Worker) GenerateChatTitle(a App) HandlerFunc {
	return func(ctx context.Context, e types.Message) error {
		generateChatTitle, ok := e.(app.GenerateChatTitle)
		if !ok {
			return fmt.Errorf("failed to cast event to generatechatTitle")
		}

		err := a.GenerateTitle(ctx, uuid.MustParse(generateChatTitle.ChatID))
		if err != nil {
			return fmt.Errorf("failed to generate title: %w", err)
		}

		return nil
	}
}

func (w *Worker) MessageAdded(a App) HandlerFunc {
	return func(ctx context.Context, e types.Message) error {
		messageAdded, ok := e.(domain.MessageAdded)
		if !ok {
			return fmt.Errorf("failed to cast event to messageadded")
		}

		err := a.ProcessAddedMessage(ctx, messageAdded)
		if err != nil {
			return fmt.Errorf("failed to handle a message added event: %w", err)
		}

		return nil
	}
}

func (w *Worker) PullOllamaModel(a App) HandlerFunc {
	return func(ctx context.Context, e types.Message) error {
		pullOllamaModel, ok := e.(app.PullOllamaModel)
		if !ok {
			return fmt.Errorf("failed to cast event to pullollamamodel")
		}

		err := a.PullOllamaModel(ctx, pullOllamaModel.Model)
		if err != nil {
			return fmt.Errorf("failed to pull ollama model: %w", err)
		}

		return nil
	}
}
