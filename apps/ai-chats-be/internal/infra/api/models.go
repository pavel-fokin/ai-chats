package api

import (
	"context"
	"log/slog"
	"net/http"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type Models interface {
	AllModelCards(context.Context) ([]domain.ModelCard, error)
}

// GetModelsLibrary returns all models cards.
func GetModelsLibrary(m Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		models, err := m.AllModelCards(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get model descriptions", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, models)
	}
}
