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

type Events interface {
	Subscribe(ctx context.Context, topic string) (chan []byte, error)
	Unsubscribe(ctx context.Context, topic string, channel chan []byte) error
}

type Message struct {
	ID     string `json:"id"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type ChatApp interface {
	CreateChat(ctx context.Context, userID uuid.UUID) (domain.Chat, error)
	AllChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error)
	AllMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error)
	SendMessage(ctx context.Context, chatID uuid.UUID, message string) (domain.Message, error)
	ChatExists(ctx context.Context, chatID uuid.UUID) (bool, error)
}

// GetChats handles the GET /api/chats endpoint.
func GetChats(chat ChatApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := MustHaveUserID(ctx)

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

		userID := MustHaveUserID(ctx)

		chat, err := chat.CreateChat(ctx, userID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create a chat", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
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
func GetEvents(app ChatApp, sse *apiutil.SSEConnections, chatEvents Events) http.HandlerFunc {
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

		conn := apiutil.NewConnection(ctx, w)
		sse.Add(conn)
		defer sse.Remove(conn)

		events, err := chatEvents.Subscribe(ctx, chatID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to subscribe to events", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}
		defer chatEvents.Unsubscribe(ctx, chatID, events)

		flusher, ok := w.(http.Flusher)
		if !ok {
			slog.ErrorContext(
				ctx,
				"failed to start the event stream",
				"err", "expected http.ResponseWriter to be an http.Flusher",
			)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		for {
			select {
			case <-conn.Closed:
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