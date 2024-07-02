package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

type AuthMock struct {
	mock.Mock
}

func (m *AuthMock) LogIn(ctx context.Context, username, password string) (domain.User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *AuthMock) SignUp(ctx context.Context, username, password string) (domain.User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestSignIn(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On("LogIn", context.Background(), "username", "password").Return(domain.User{}, nil)

		LogIn(auth)(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "LogIn", 1)
	})

	t.Run("Failure", func(t *testing.T) {
		// Setup.
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On(
			"LogIn",
			context.Background(), "username", "password",
		).Return(
			domain.User{}, fmt.Errorf("some error"),
		)

		// Test.
		LogIn(auth)(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "LogIn", 1)
	})

	t.Run("Invalid request", func(t *testing.T) {
		// Setup.
		body := `{"username": "username"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}

		// Test.
		LogIn(auth)(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "LogIn", 0)
	})
}
