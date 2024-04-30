package server

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bot/internal/app"
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
		var req PostMessagesRequest
		if err := ParseJSON(r, &req); err != nil {
			AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		answer, err := chat.SendMessage(r.Context(), uuid.Nil, uuid.Nil, req.Message.Text)
		if err != nil {
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
