package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app"
	"pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type Chat interface {
	CreateChat(ctx context.Context) (domain.Chat, error)
	SendMessage(ctx context.Context, chatId uuid.UUID, message string) (app.Message, error)
}

// PostChats handles the POST /api/chats endpoint.
func PostChats(chat Chat) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chat, err := chat.CreateChat(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create a chat", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, chat, http.StatusOK)
	}
}

// PostMessages handles the POST /api/chats/{uuid}/messages endpoint.
func PostMessages(chat Chat) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		var req PostMessagesRequest
		if err := apiutil.ParseJSON(r, &req); err != nil {
			apiutil.AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		answer, err := chat.SendMessage(ctx, uuid.MustParse(chatID), req.Message.Text)
		if err != nil {
			slog.ErrorContext(ctx, "failed to send a message", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, PostMessagesResponse{
			Message{
				Sender: "AI",
				Text:   answer.Text,
			},
		}, http.StatusOK)
	}
}
