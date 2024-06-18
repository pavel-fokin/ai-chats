package api

import (
	"context"
	"log/slog"
	"net/http"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"
)

type OllamaApp interface {
	AllOllamaModels(ctx context.Context) ([]domain.Model, error)
}

// GetOllamaModels handles the GET /api/ollama-models endpoint.
func GetOllamaModels(ollama OllamaApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		models, err := ollama.AllOllamaModels(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get ollama models", "err", err)
			apiutil.AsErrorResponse(w, ErrInternal, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, NewGetOllamaModelsResponse(models), http.StatusOK)
	}
}
