package server

import (
	"net/http"

	"pavel-fokin/ai/apps/ai-bot/internal/app"
)

type PostMessagesRequest struct {
	app.Message
}

type PostMessagesResponse struct {
	app.Message
}

// PostMessages handles the POST /api/messages endpoint.
func PostMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req PostMessagesRequest
		if err := ParseJSON(r, &req); err != nil {
			AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		message, err := app.SendMessage(r.Context(), req.Message.Text)
		if err != nil {
			AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, PostMessagesResponse{Message: message}, http.StatusOK)
	}
}
