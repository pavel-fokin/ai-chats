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
	ID     string `json:"id"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type ChatApp interface {
	CreateChat(ctx context.Context) (domain.Chat, error)
	AllChats(ctx context.Context) ([]domain.Chat, error)
	AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error)
	SendMessage(ctx context.Context, chatId uuid.UUID, message string) (app.Message, error)
}

// GetChats handles the GET /api/chats endpoint.
func GetChats(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chats, err := chat.AllChats(ctx)
		slog.InfoContext(ctx, "get chats", "chats", chats)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get chats", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, NewGetChatsResponse(chats), http.StatusOK)
	}
}

// PostChats handles the POST /api/chats endpoint.
func PostChats(chat ChatApp) http.HandlerFunc {
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

// GetMessages handles the GET /api/chats/{uuid}/messages endpoint.
func GetMessages(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		messages, err := chat.AllMessages(ctx, uuid.MustParse(chatID))
		if err != nil {
			slog.ErrorContext(ctx, "failed to get messages", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, NewGetMessagesResponse(messages), http.StatusOK)
	}
}

// PostMessages handles the POST /api/chats/{uuid}/messages endpoint.
func PostMessages(chat ChatApp) http.HandlerFunc {
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
