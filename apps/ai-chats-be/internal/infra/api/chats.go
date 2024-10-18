package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"ai-chats/internal/domain"
	"ai-chats/internal/pkg/types"
)

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (chan types.Message, error)
	Unsubscribe(ctx context.Context, topic string, channel chan types.Message) error
}

type Chats interface {
	ChatExists(ctx context.Context, chatID domain.ChatID) (bool, error)
	ChatMessages(ctx context.Context, chatID domain.ChatID) ([]domain.Message, error)
	CreateChat(ctx context.Context, userID domain.UserID, defaultModel, message string) (domain.Chat, error)
	DeleteChat(ctx context.Context, chatID domain.ChatID) error
	FindChatByID(ctx context.Context, chatID domain.ChatID) (domain.Chat, error)
	FindChatsByUserID(ctx context.Context, userID domain.UserID) ([]domain.Chat, error)
	GenerateChatTitleAsync(ctx context.Context, chatID domain.ChatID) error
	SendMessage(ctx context.Context, userID domain.UserID, chatID domain.ChatID, message string) (domain.Message, error)
}

// GetChats handles the GET /api/chats endpoint.
func GetChats(app Chats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := MustHaveUserID(ctx)

		chats, err := app.FindChatsByUserID(ctx, userID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get chats", "userID", userID, "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, NewGetChatsResponse(chats))
	}
}

// PostChats handles the POST /api/chats endpoint.
func PostChats(app Chats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := MustHaveUserID(ctx)

		defaultModel := ""
		message := ""
		if r.Body != nil {
			var req PostChatsRequest
			if err := ParseJSON(r, &req); err != nil {
				slog.ErrorContext(ctx, "failed to parse the request", "err", err)
				WriteErrorResponse(w, http.StatusBadRequest, BadRequest)
				return
			}
			defaultModel = req.DefaultModel
			message = req.Message
		}

		chat, err := app.CreateChat(ctx, userID, defaultModel, message)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create a chat", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, NewPostChatsResponse(chat))
	}
}

// GetChat handles the GET /api/chats/{uuid} endpoint.
func GetChat(app Chats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		chat, err := app.FindChatByID(ctx, uuid.MustParse(chatID))
		if err != nil {
			slog.ErrorContext(ctx, "failed to get a chat", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, NewGetChatResponse(chat))
	}
}

// DeleteChat handles the DELETE /api/chats/{uuid} endpoint.
func DeleteChat(app Chats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		if err := app.DeleteChat(ctx, uuid.MustParse(chatID)); err != nil {
			switch err {
			case domain.ErrChatNotFound:
				slog.ErrorContext(ctx, "chat not found", "chatID", chatID)
				WriteErrorResponse(w, http.StatusNotFound, NotFound)
			default:
				slog.ErrorContext(ctx, "failed to delete a chat", "err", err)
				WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			}
			return
		}

		WriteSuccessResponse(w, http.StatusNoContent, nil)
	}
}

// GetMessages handles the GET /api/chats/{uuid}/messages endpoint.
func GetMessages(app Chats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		messages, err := app.ChatMessages(ctx, uuid.MustParse(chatID))
		if err != nil {
			switch err {
			case domain.ErrChatNotFound:
				slog.ErrorContext(ctx, "chat not found", "chatID", chatID)
				WriteErrorResponse(w, http.StatusNotFound, NotFound)
			default:
				slog.ErrorContext(ctx, "failed to get messages", "err", err)
				WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			}
			return
		}

		WriteSuccessResponse(w, http.StatusOK, NewGetMessagesResponse(messages))
	}
}

// PostMessages handles the POST /api/chats/{uuid}/messages endpoint.
func PostMessages(app Chats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")
		userID := MustHaveUserID(ctx)

		var req PostMessagesRequest
		if err := ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse the request", "err", err)
			WriteErrorResponse(w, http.StatusBadRequest, BadRequest)
			return
		}

		_, err := app.SendMessage(ctx, userID, uuid.MustParse(chatID), req.Text)
		if err != nil {
			slog.ErrorContext(ctx, "failed to send a message", "chatID", chatID, "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusNoContent, nil)
	}
}

// PostGenerateChatTitle handles the POST /api/chats/{uuid}/generate-title endpoint.
func PostGenerateChatTitle(app Chats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		err := app.GenerateChatTitleAsync(ctx, uuid.MustParse(chatID))
		if err != nil {
			slog.ErrorContext(ctx, "failed to generate a chat title", "chatID", chatID, "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusNoContent, nil)
	}
}

// GetChatEvents handles the GET /api/chats/{uuid}/events endpoint.
func GetChatEvents(app Chats, sse *SSEConnections, subscriber Subscriber) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		chatID := chi.URLParam(r, "uuid")

		chatExists, err := app.ChatExists(ctx, uuid.MustParse(chatID))
		if err != nil {
			slog.ErrorContext(ctx, "failed to check if the chat exists", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}
		if !chatExists {
			slog.ErrorContext(ctx, "chat not found", "chatID", chatID)
			WriteErrorResponse(w, http.StatusNotFound, NotFound)
			return
		}

		conn := sse.AddConnection()
		defer sse.Remove(conn)

		events, err := subscriber.Subscribe(ctx, chatID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to subscribe to events", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}
		defer subscriber.Unsubscribe(ctx, chatID, events)

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher := w.(http.Flusher)
		for {
			select {
			case <-ctx.Done():
				return
			case <-conn.Closed:
				return
			case event := <-events:
				if err := WriteServerSentEvent(w, event); err != nil {
					slog.ErrorContext(ctx, "failed to write an event", "err", err)
					return
				}
				flusher.Flush()
			}
		}
	}
}
