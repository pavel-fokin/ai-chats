package api

import (
	"context"
	"log/slog"
	"net/http"

	"ai-chats/internal/domain"
)

type Models interface {
	AllModelCards(context.Context) ([]domain.ModelCard, error)
}

type GetModelsLibraryResponse struct {
	ModelCards []domain.ModelCard `json:"modelCards"`
}

func NewGetModelsLibraryResponse(modelCards []domain.ModelCard) GetModelsLibraryResponse {
	return GetModelsLibraryResponse{
		ModelCards: modelCards,
	}
}

// GetModelsLibrary handles the GET /api/models/library endpoint.
func GetModelsLibrary(m Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		models, err := m.AllModelCards(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get model descriptions", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, NewGetModelsLibraryResponse(models))
	}
}
