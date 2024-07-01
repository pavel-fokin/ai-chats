package api

import (
	"context"
	"log/slog"
	"net/http"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"

	"github.com/go-chi/chi/v5"
)

type OllamaApp interface {
	ListModels(context.Context) ([]domain.Model, error)
	PullModel(context.Context, string) error
	DeleteModel(context.Context, string) error
}

// GetOllamaModels handles the GET /api/ollama-models endpoint.
func GetOllamaModels(ollama OllamaApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		models, err := ollama.ListModels(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get ollama models", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, NewGetOllamaModelsResponse(models), http.StatusOK)
	}
}

// PostOllamaModels handles the POST /api/ollama-models endpoint.
func PostOllamaModels(ollama OllamaApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req PostOllamaModelsRequest
		if err := apiutil.ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse the request", "err", err)
			apiutil.AsErrorResponse(w, ErrBadRequest, http.StatusBadRequest)
			return
		}

		err := ollama.PullModel(ctx, req.Model)
		if err != nil {
			slog.ErrorContext(ctx, "failed to pull ollama model", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, nil, http.StatusNoContent)
	}
}

// DeleteOllamaModels handles the DELETE /api/ollama-models/{model} endpoint.
func DeleteOllamaModel(ollama OllamaApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		model := chi.URLParam(r, "model")

		err := ollama.DeleteModel(ctx, model)
		if err != nil {
			slog.ErrorContext(ctx, "failed to delete ollama model", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, nil, http.StatusNoContent)
	}
}
