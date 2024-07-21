package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"ai-chats/internal/domain"
)

type Auth interface {
	LogIn(ctx context.Context, username, password string) (domain.User, error)
	SignUp(ctx context.Context, username, password string) (domain.User, error)
}

// LogIn logs in a user.
func LogIn(app Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req LogInRequest
		if err := ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse request body", "err", err)
			WriteErrorResponse(w, http.StatusBadRequest, BadRequest)
			return
		}

		user, err := app.LogIn(ctx, req.Username, req.Password)
		if err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("failed to sign in user - %s", req.Username), "err", err)
			switch {
			case errors.Is(err, domain.ErrUserNotFound):
				WriteErrorResponse(w, http.StatusUnauthorized, UsernameOrPasswordIsIncorrect)
			case errors.Is(err, domain.ErrInvalidPassword):
				WriteErrorResponse(w, http.StatusUnauthorized, UsernameOrPasswordIsIncorrect)
			default:
				WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			}
			return
		}

		accessToken, err := NewAccessToken(user.ID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create access token", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, LogInResponse{accessToken})
	}
}

// SignUp signs up a user.
func SignUp(app Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req SignUpRequest
		if err := ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse request body", "err", err)
			WriteErrorResponse(w, http.StatusBadRequest, BadRequest)
			return
		}

		user, err := app.SignUp(ctx, req.Username, req.Password)
		if err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("failed to add user - %s", req.Username), "err", err)
			switch {
			case errors.Is(err, domain.ErrUserAlreadyExists):
				WriteErrorResponse(w, http.StatusConflict, UsernameIsTaken)
			default:
				WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			}
			return
		}

		accessToken, err := NewAccessToken(user.ID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create access token", "err", err)
			WriteErrorResponse(w, http.StatusInternalServerError, InternalError)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, SignUpResponse{accessToken})
	}
}
