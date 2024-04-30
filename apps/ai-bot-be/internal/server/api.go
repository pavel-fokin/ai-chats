package server

import (
	"context"
	"net/http"

	"pavel-fokin/ai/apps/ai-bot/internal/app"
)

type PostMessagesRequest struct {
	app.Message
}

type PostMessagesResponse struct {
	app.Message
}

type Chat interface {
	SendMessage(ctx context.Context, chatId, message string) (app.Message, error)
}

// PostMessages handles the POST /api/messages endpoint.
func PostMessages(chat Chat) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostMessagesRequest
		if err := ParseJSON(r, &req); err != nil {
			AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		message, err := chat.SendMessage(r.Context(), "id", req.Message.Text)
		if err != nil {
			AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, PostMessagesResponse{Message: message}, http.StatusOK)
	}
}
