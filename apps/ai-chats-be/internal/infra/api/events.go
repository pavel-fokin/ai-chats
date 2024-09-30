package api

import (
	"log/slog"
	"net/http"
)

// GetAppEvents handles the GET /api/events/app endpoint.
func GetAppEvents(app Chats, sse *SSEConnections, subscriber Subscriber) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := MustHaveUserID(ctx)

		conn := sse.AddConnection()
		defer sse.Remove(conn)

		events, err := subscriber.Subscribe(ctx, userID.String())
		if err != nil {
			slog.ErrorContext(ctx, "failed to subscribe to events", "err", err)
			AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}
		defer subscriber.Unsubscribe(ctx, userID.String(), events)

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Channel to signal when the handler should exit
		done := make(chan struct{})

		// Listen for shutdown signal in a separate goroutine
		go func() {
			select {
			case <-ctx.Done():
				// Client disconnected
				close(done)
			case <-conn.Closed:
				// Server is shutting down
				close(done)
			}
		}()

		flusher := w.(http.Flusher)
		for {
			select {
			case <-done:
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
