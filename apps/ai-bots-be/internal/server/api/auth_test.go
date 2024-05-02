package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pavel-fokin/ai/apps/ai-bots-be/internal/app"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AuthMock struct {
	mock.Mock
}

func (m *AuthMock) SignIn(ctx context.Context, username, password string) (*app.User, error) {
	args := m.Called(ctx, username, password)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*app.User), args.Error(1)
}

func (m *AuthMock) SignUp(ctx context.Context, username, password string) (*app.User, error) {
	args := m.Called(ctx, username, password)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*app.User), args.Error(1)
}

func TestSignIn(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/signin", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On("SignIn", context.Background(), "username", "password").Return(&app.User{}, nil)

		SignIn(auth, "token")(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignIn", 1)
	})

	t.Run("Failure", func(t *testing.T) {
		// Setup.
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/signup", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On("SignIn", context.Background(), "username", "password").Return(nil, fmt.Errorf("some error"))

		// Test.
		SignIn(auth, "token")(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignIn", 1)
	})

	t.Run("Invalid request", func(t *testing.T) {
		// Setup.
		body := `{"username": "username"}`
		req, _ := http.NewRequest("POST", "/signin", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}

		// Test.
		SignIn(auth, "token")(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignIn", 0)
	})
}
