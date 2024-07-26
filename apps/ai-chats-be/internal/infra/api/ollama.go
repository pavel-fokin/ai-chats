package api

import (
	"context"
	"log/slog"
	"net/http"

	"ai-chats/internal/domain"

	"github.com/go-chi/chi/v5"
)

type Ollama interface {
	ListOllamaModels(context.Context) ([]domain.OllamaModel, error)
	PullOllamaModel(context.Context, string) error
	DeleteOllamaModel(context.Context, string) error
}

// GetOllamaModels handles the GET /api/ollama/models endpoint.
func GetOllamaModels(app Ollama) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		models, err := app.ListOllamaModels(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get ollama models", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, NewGetOllamaModelsResponse(models))
	}
}

// PostOllamaModels handles the POST /api/ollama/models endpoint.
func PostOllamaModels(app Ollama) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req PostOllamaModelsRequest
		if err := ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse the request", "err", err)
			AsErrorResponse(w, ErrBadRequest, http.StatusBadRequest)
			return
		}

		err := app.PullOllamaModel(ctx, req.Model)
		if err != nil {
			slog.ErrorContext(ctx, "failed to pull ollama model", "err", err)
			AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, nil, http.StatusNoContent)
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
			AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, nil, http.StatusNoContent)
	}
}
