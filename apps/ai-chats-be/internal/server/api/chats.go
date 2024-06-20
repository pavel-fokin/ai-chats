package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"
)

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (chan []byte, error)
	Unsubscribe(ctx context.Context, topic string, channel chan []byte) error
}

type ChatApp interface {
	CreateChat(ctx context.Context, userID uuid.UUID, message string) (domain.Chat, error)
	DeleteChat(ctx context.Context, chatID domain.ChatID) error
	FindChatByID(ctx context.Context, chatID domain.ChatID) (domain.Chat, error)
	AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error)
	AllMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error)
	SendMessage(ctx context.Context, chatID domain.ChatID, message string) (domain.Message, error)
	ChatExists(ctx context.Context, chatID domain.ChatID) (bool, error)
}

// GetChats handles the GET /api/chats endpoint.
func GetChats(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := apiutil.MustHaveUserID(ctx)

		chats, err := chat.AllChats(ctx, userID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get chats", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, NewGetChatsResponse(chats), http.StatusOK)
	}
}

// PostChats handles the POST /api/chats endpoint.
func PostChats(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := apiutil.MustHaveUserID(ctx)

		message := ""
		if r.Body != nil {
			var req PostChatsRequest
			if err := apiutil.ParseJSON(r, &req); err != nil {
				slog.ErrorContext(ctx, "failed to parse the request", "err", err)
				apiutil.AsErrorResponse(w, ErrBadRequest, http.StatusBadRequest)
				return
			}
			message = req.Message
		}

		chat, err := chat.CreateChat(ctx, userID, message)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create a chat", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, NewPostChatsResponse(chat), http.StatusOK)
	}
}

// GetChat handles the GET /api/chats/{uuid} endpoint.
func GetChat(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		chat, err := chat.FindChatByID(ctx, uuid.MustParse(chatID))
		if err != nil {
			slog.ErrorContext(ctx, "failed to get a chat", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, NewGetChatResponse(chat), http.StatusOK)
	}
}

// DeleteChat handles the DELETE /api/chats/{uuid} endpoint.
func DeleteChat(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		if err := chat.DeleteChat(ctx, uuid.MustParse(chatID)); err != nil {
			switch err {
			case domain.ErrChatNotFound:
				apiutil.AsErrorResponse(w, ErrNotFound, http.StatusNotFound)
				return
			default:
				slog.ErrorContext(ctx, "failed to delete a chat", "err", err)
				apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
				return
			}
		}

		apiutil.AsSuccessResponse(w, nil, http.StatusNoContent)
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
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
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
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusBadRequest)
			return
		}

		answer, err := chat.SendMessage(ctx, uuid.MustParse(chatID), req.Message.Text)
		if err != nil {
			slog.ErrorContext(ctx, "failed to send a message", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
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

// GetEvents handles the GET /api/chats/{uuid}/events endpoint.
func GetEvents(app ChatApp, sse *apiutil.SSEConnections, subscriber Subscriber) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		if exists, err := app.ChatExists(ctx, uuid.MustParse(chatID)); err != nil {
			slog.ErrorContext(ctx, "failed to check if the chat exists", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		} else if !exists {
			apiutil.AsErrorResponse(w, ErrNotFound, http.StatusNotFound)
			return
		}

		conn := sse.AddConnection()
		defer sse.Remove(conn)

		events, err := subscriber.Subscribe(ctx, chatID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to subscribe to events", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}
		defer subscriber.Unsubscribe(ctx, chatID, events)

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher := w.(http.Flusher)
		for {
			select {
			case <-conn.Closed:
				apiutil.AsSuccessResponse(w, nil, http.StatusNoContent)
				return
			case <-ctx.Done():
				return
			case event := <-events:
				if err := apiutil.WriteEvent(w, event); err != nil {
					slog.ErrorContext(ctx, "failed to write an event", "err", err)
					return
				}
				flusher.Flush()
			}
		}
	}
}
