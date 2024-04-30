package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app"
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type PostMessagesRequest struct {
	Message
}

type PostMessagesResponse struct {
	Message
}

type Chat interface {
	SendMessage(ctx context.Context, userID uuid.UUID, chatId uuid.UUID, message string) (app.Message, error)
}

// PostMessages handles the POST /api/chats/{uuid}/messages endpoint.
func PostChatMessages(chat Chat) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req PostMessagesRequest
		if err := ParseJSON(r, &req); err != nil {
			AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		answer, err := chat.SendMessage(ctx, uuid.Nil, uuid.Nil, req.Message.Text)
		if err != nil {
			slog.ErrorContext(ctx, "failed to send a message", "err", err)
			AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, PostMessagesResponse{
			Message{
				Sender: "AI",
				Text:   answer.Text,
			},
		}, http.StatusOK)
	}
}
