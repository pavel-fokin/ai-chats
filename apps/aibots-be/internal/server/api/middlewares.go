package api

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"

	"github.com/google/uuid"
)

type UserID string

// UserIDFromContext returns the UserID value from the context.
func MustHaveUserID(ctx context.Context) uuid.UUID {
	v := ctx.Value(UserID("UserID"))
	if v == nil {
		panic("missing user ID")
	}

	userID, ok := v.(uuid.UUID)
	if !ok {
		panic("invalid user ID")
	}

	return userID
}

// AuthToken is a middleware that checks for the presence of an Authorization header.
func AuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			slog.ErrorContext(r.Context(), "missing auth token")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		accessToken := strings.TrimPrefix(authToken, "Bearer ")
		claims, err := apiutil.VerifyAccessToken(accessToken)
		if err != nil {
			slog.ErrorContext(r.Context(), "failed to verify access token", "err", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserID("UserID"), claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
