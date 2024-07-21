package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"ai-chats/internal/domain"
	"ai-chats/internal/server"
)

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (chan []byte, error)
	Unsubscribe(ctx context.Context, topic string, channel chan []byte) error
}

type ChatApp interface {
	AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error)
	AllMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error)
	ChatExists(ctx context.Context, chatID domain.ChatID) (bool, error)
	CreateChat(ctx context.Context, userID uuid.UUID, defaultModel, message string) (domain.Chat, error)
	DeleteChat(ctx context.Context, chatID domain.ChatID) error
	FindChatByID(ctx context.Context, chatID domain.ChatID) (domain.Chat, error)
	SendMessage(ctx context.Context, userID domain.UserID, chatID domain.ChatID, message string) (domain.Message, error)
}

// GetChats handles the GET /api/chats endpoint.
func GetChats(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := MustHaveUserID(ctx)

		chats, err := chat.AllChats(ctx, userID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get chats", "userID", userID, "err", err)
			AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, NewGetChatsResponse(chats), http.StatusOK)
	}
}

// PostChats handles the POST /api/chats endpoint.
func PostChats(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := MustHaveUserID(ctx)

		defaultModel := ""
		message := ""
		if r.Body != nil {
			var req PostChatsRequest
			if err := ParseJSON(r, &req); err != nil {
				slog.ErrorContext(ctx, "failed to parse the request", "err", err)
				AsErrorResponse(w, ErrBadRequest, http.StatusBadRequest)
				return
			}
			defaultModel = req.DefaultModel
			message = req.Message
		}

		chat, err := chat.CreateChat(ctx, userID, defaultModel, message)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create a chat", "err", err)
			AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, NewPostChatsResponse(chat), http.StatusOK)
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
			AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, NewGetChatResponse(chat), http.StatusOK)
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
				slog.ErrorContext(ctx, "chat not found", "chatID", chatID)
				AsErrorResponse(w, ErrNotFound, http.StatusNotFound)
			default:
				slog.ErrorContext(ctx, "failed to delete a chat", "err", err)
				AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			}
			return
		}

		AsSuccessResponse(w, nil, http.StatusNoContent)
	}
}

// GetMessages handles the GET /api/chats/{uuid}/messages endpoint.
func GetMessages(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		messages, err := chat.AllMessages(ctx, uuid.MustParse(chatID))
		if err != nil {
			switch err {
			case domain.ErrChatNotFound:
				slog.ErrorContext(ctx, "chat not found", "chatID", chatID)
				AsErrorResponse(w, ErrNotFound, http.StatusNotFound)
			default:
				slog.ErrorContext(ctx, "failed to get messages", "err", err)
				AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			}
			return
		}

		AsSuccessResponse(w, NewGetMessagesResponse(messages), http.StatusOK)
	}
}

// PostMessages handles the POST /api/chats/{uuid}/messages endpoint.
func PostMessages(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")
		userID := MustHaveUserID(ctx)

		var req PostMessagesRequest
		if err := ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse the request", "err", err)
			AsErrorResponse(w, ErrInternal, http.StatusBadRequest)
			return
		}

		_, err := chat.SendMessage(ctx, userID, uuid.MustParse(chatID), req.Text)
		if err != nil {
			slog.ErrorContext(ctx, "failed to send a message", "chatID", chatID, "err", err)
			AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, nil, http.StatusNoContent)
	}
}

// GetEvents handles the GET /api/chats/{uuid}/events endpoint.
func GetEvents(app ChatApp, sse *server.SSEConnections, subscriber Subscriber) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		chatExists, err := app.ChatExists(ctx, uuid.MustParse(chatID))
		if err != nil {
			slog.ErrorContext(ctx, "failed to check if the chat exists", "err", err)
			AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}
		if !chatExists {
			slog.ErrorContext(ctx, "chat not found", "chatID", chatID)
			AsErrorResponse(w, ErrNotFound, http.StatusNotFound)
			return
		}

		conn := sse.AddConnection()
		defer sse.Remove(conn)

		events, err := subscriber.Subscribe(ctx, chatID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to subscribe to events", "err", err)
			AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
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
				AsSuccessResponse(w, nil, http.StatusNoContent)
				return
			case <-ctx.Done():
				return
			case event := <-events:
				if err := WriteEvent(w, event); err != nil {
					slog.ErrorContext(ctx, "failed to write an event", "err", err)
					return
				}
				flusher.Flush()
			}
		}
	}
}
