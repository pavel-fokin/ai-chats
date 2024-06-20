package apiutil

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthToken(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(UserIDCtxKey)
		if userID == nil {
			t.Error("UserID context value is not set")
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	accessToken, err := NewAccessToken(uuid.New())
	assert.NoError(t, err)

	// Create a test request with a valid access token.
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	res := httptest.NewRecorder()

	AuthHeader(handler).ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestMustHaveUserID(t *testing.T) {
	t.Run("Must have UserID", func(t *testing.T) {
		userID := uuid.New()
		ctx := context.WithValue(context.Background(), UserIDCtxKey, userID)

		gotUserID := MustHaveUserID(ctx)

		assert.Equal(t, userID, gotUserID)
	})

	t.Run("Missing UserID", func(t *testing.T) {
		assert.Panics(t, func() {
			MustHaveUserID(context.Background())
		})
	})
}

func TestAuthParam(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := r.Context().Value(UserIDCtxKey)
			if userID == nil {
				t.Error("UserID context value is not set")
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		accessToken, err := NewAccessToken(uuid.New())
		assert.NoError(t, err)

		// Create a test request with a valid access token.
		req := httptest.NewRequest(http.MethodGet, "/?accessToken="+accessToken, nil)
		res := httptest.NewRecorder()

		// Call the AuthParam middleware.
		AuthParam(handler).ServeHTTP(res, req)

		// Check the response status code.
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Missing AccessToken", func(t *testing.T) {
		// Create a test handler.
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Create a test request without an access token.
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		// Call the AuthParam middleware.
		AuthParam(handler).ServeHTTP(res, req)

		// Check the response status code.
		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}
