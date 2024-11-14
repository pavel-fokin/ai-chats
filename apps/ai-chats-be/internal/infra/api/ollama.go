package api

import (
	"context"
	"log/slog"
	"net/http"

	"ai-chats/internal/domain"

	"github.com/go-chi/chi/v5"
)

type Ollama interface {
	GetOllamaModelsLibrary(context.Context) ([]*domain.ModelCard, error)
	DeleteOllamaModel(context.Context, string) error
	FindOllamaModels(context.Context, domain.OllamaModelsFilter) ([]domain.OllamaModel, error)
	PullOllamaModelAsync(context.Context, string) error
}

// GetOllamaModels handles the GET /api/ollama/models endpoint.
func GetOllamaModels(app Ollama) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		filter, err := ParseOllamaModelsQuery(r.URL.Query().Encode())
		if err != nil {
			slog.ErrorContext(ctx, "failed to parse ollama models query", "err", err)
			WriteErrorResponse(w, http.StatusBadRequest, BadRequest)
			return
		}

		models, err := app.FindOllamaModels(ctx, filter)
		if err != nil {
			slog.ErrorContext(ctx, "failed to find ollama models", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, NewGetOllamaModelsResponse(models))
	}
}

// GetOllamaModelsLibrary handles the GET /api/ollama/models-library endpoint.
func GetOllamaModelsLibrary(app Ollama) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		models, err := app.GetOllamaModelsLibrary(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get ollama models library", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, NewGetOllamaModelsLibraryResponse(models))
	}
}

// PostOllamaModels handles the POST /api/ollama/models endpoint.
func PostOllamaModels(app Ollama) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req PostOllamaModelsRequest
		if err := ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse the request", "err", err)
			WriteErrorResponse(w, http.StatusBadRequest, BadRequest)
			return
		}

		err := app.PullOllamaModelAsync(ctx, req.Model)
		if err != nil {
			slog.ErrorContext(ctx, "failed to pull ollama model", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusNoContent, nil)
	}
}

// DeleteOllamaModels handles the DELETE /api/ollama/models/{model} endpoint.
func DeleteOllamaModel(app Ollama) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		model := chi.URLParam(r, "model")

		err := app.DeleteOllamaModel(ctx, model)
		if err != nil {
			slog.ErrorContext(ctx, "failed to delete ollama model", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusNoContent, nil)
	}
}

// GetOllamaModelPullingEvents handles the GET /api/ollama/models/{model}/pulling-events endpoint.
func GetOllamaModelPullingEvents(app Ollama, sse *SSEConnections, subscriber Subscriber) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		conn := sse.AddConnection()
		defer sse.Remove(conn)

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		model := chi.URLParam(r, "model")

		events, err := subscriber.Subscribe(ctx, model)
		if err != nil {
			slog.ErrorContext(ctx, "failed to subscribe to events", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}
		defer subscriber.Unsubscribe(ctx, model, events)

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
