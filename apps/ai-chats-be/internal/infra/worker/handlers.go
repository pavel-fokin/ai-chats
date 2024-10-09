package worker

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	appPkg "ai-chats/internal/app"
	"ai-chats/internal/app/commands"
	"ai-chats/internal/domain/events"
	"ai-chats/internal/pkg/types"
)

func (w *Worker) SetupHandlers(app App) {
	w.RegisterHandler(appPkg.GenerateChatTitleTopic, 1, w.GenerateChatTitle(app))
	w.RegisterHandler(appPkg.MessageAddedTopic, 1, w.MessageAdded(app))
	w.RegisterHandler(appPkg.PullOllamaModelTopic, 1, w.PullOllamaModel(app))
}

func (w *Worker) GenerateChatTitle(app App) HandlerFunc {
	return func(ctx context.Context, e types.Message) error {
		generateChatTitle, ok := e.(commands.GenerateChatTitle)
		if !ok {
			return fmt.Errorf("failed to cast event to generatechatTitle")
		}

		err := app.GenerateTitle(ctx, uuid.MustParse(generateChatTitle.ChatID))
		if err != nil {
			return fmt.Errorf("failed to generate title: %w", err)
		}

		return nil
	}
}

func (w *Worker) MessageAdded(app App) HandlerFunc {
	return func(ctx context.Context, e types.Message) error {
		messageAdded, ok := e.(events.MessageAdded)
		if !ok {
			return fmt.Errorf("failed to cast event to messageadded")
		}

		err := app.ProcessAddedMessage(ctx, messageAdded)
		if err != nil {
			return fmt.Errorf("failed to handle a message added event: %w", err)
		}

		return nil
	}
}

func (w *Worker) PullOllamaModel(app App) HandlerFunc {
	return func(ctx context.Context, e types.Message) error {
		pullOllamaModel, ok := e.(commands.PullOllamaModel)
		if !ok {
			return fmt.Errorf("failed to cast event to pullollamamodel")
		}

		err := app.PullOllamaModel(ctx, pullOllamaModel.Model)
		if err != nil {
			return fmt.Errorf("failed to pull ollama model: %w", err)
		}

		return nil
	}
}
