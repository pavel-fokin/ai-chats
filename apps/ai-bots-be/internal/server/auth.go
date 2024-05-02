package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app"
)

type Auth interface {
	SignIn(ctx context.Context, username, password string) (*app.User, error)
	SignUp(ctx context.Context, username, password string) (*app.User, error)
}

// SignIn signs in a user.
func SignIn(app Auth, tokenSigningKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req SignInRequest
		if err := ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse request body", "err", err)
			AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		user, err := app.SignIn(ctx, req.Username, req.Password)
		if err != nil {
			slog.ErrorContext(ctx, "failed to sign in user", "err", err)
			AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		accessToken, err := NewAccessToken(user.ID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create access token", "err", err)
			AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, &SignInResponse{
			AccessToken: accessToken,
		}, http.StatusOK)
	}
}

// SignUp signs up a user.
func SignUp(app Auth, tokenSigningKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req SignUpRequest
		if err := ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse request body", "err", err)
			AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		user, err := app.SignUp(ctx, req.Username, req.Password)
		if err != nil {
			slog.ErrorContext(ctx, "failed to sign up user", "err", err)
			http.Error(w, fmt.Sprintf("failed to sign up user: %v", err), http.StatusInternalServerError)
			return
		}

		accessToken, err := NewAccessToken(user.ID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create access token", "err", err)
			AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		AsSuccessResponse(w, SignInResponse{
			AccessToken: accessToken,
		}, http.StatusOK)
	}
}
