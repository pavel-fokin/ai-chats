package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthToken(t *testing.T) {
	// Create a test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the UserID context value is set correctly
		userID := r.Context().Value(UserID("UserID"))
		if userID == nil {
			t.Error("UserID context value is not set")
			return
		}

		// Write a response
		w.WriteHeader(http.StatusOK)
	})

	accessToken, err := apiutil.NewAccessToken(uuid.New())
	assert.NoError(t, err)

	// Create a test request with a valid access token
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// Create a test response recorder
	res := httptest.NewRecorder()

	// Call the AuthToken middleware
	AuthToken(handler).ServeHTTP(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, res.Code)
}
