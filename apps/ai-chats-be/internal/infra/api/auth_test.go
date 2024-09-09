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

	"ai-chats/internal/domain"
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

func TestAuthAPILogIn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On("LogIn", context.Background(), "username", "password").Return(domain.User{}, nil)

		LogIn(auth, "signingKey")(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "LogIn", 1)
	})

	t.Run("failure", func(t *testing.T) {
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

		LogIn(auth, "signingKey")(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "LogIn", 1)
	})

	t.Run("invalid request", func(t *testing.T) {
		body := `{"username": "username"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}

		LogIn(auth, "signingKey")(w, req)

		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "LogIn", 0)
	})

	t.Run("username is incorrect", func(t *testing.T) {
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On(
			"LogIn",
			context.Background(), "username", "password",
		).Return(
			domain.User{}, domain.ErrUserNotFound,
		)

		LogIn(auth, "signingKey")(w, req)

		resp := w.Result()
		assert.Equal(t, 401, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "LogIn", 1)
	})

	t.Run("password is incorrect", func(t *testing.T) {
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On(
			"LogIn",
			context.Background(), "username", "password",
		).Return(
			domain.User{}, domain.ErrInvalidPassword,
		)

		LogIn(auth, "signingKey")(w, req)

		resp := w.Result()
		assert.Equal(t, 401, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "LogIn", 1)
	})
}

func TestAuthAPISignUP(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/signup", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On("SignUp", context.Background(), "username", "password").Return(domain.User{}, nil)

		SignUp(auth, "signingKey")(w, req)

		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignUp", 1)
	})

	t.Run("failure", func(t *testing.T) {
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/signup", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On(
			"SignUp",
			context.Background(), "username", "password",
		).Return(
			domain.User{}, fmt.Errorf("some error"),
		)

		SignUp(auth, "signingKey")(w, req)

		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignUp", 1)
	})

	t.Run("invalid request", func(t *testing.T) {
		body := `{"username": "username"}`
		req, _ := http.NewRequest("POST", "/signup", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}

		SignUp(auth, "signingKey")(w, req)

		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignUp", 0)
	})

	t.Run("username is taken", func(t *testing.T) {
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/signup", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On(
			"SignUp",
			context.Background(), "username", "password",
		).Return(
			domain.User{}, domain.ErrUserAlreadyExists,
		)

		SignUp(auth, "signingKey")(w, req)

		resp := w.Result()
		assert.Equal(t, 409, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignUp", 1)
	})
}
