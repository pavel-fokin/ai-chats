package app

import (
	"context"
	"errors"
	"pavel-fokin/ai/apps/ai-bots-be/internal/app/apputil"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUsers struct {
	mock.Mock
}

func (m *MockUsers) AddUser(ctx context.Context, user domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUsers) FindUser(ctx context.Context, username string) (domain.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUsers) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestSignUp(t *testing.T) {
	mockUsers := &MockUsers{}
	mockUsers.On("AddUser", context.Background(), mock.Anything).Return(nil)

	app := &App{
		users: mockUsers,
	}

	user, err := app.SignUp(context.Background(), "username", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	// Verify that password is not stored in plain text.
	assert.NotEqual(t, "password", user.PasswordHash)
}

func TestSignIn(t *testing.T) {
	t.Run("user not found", func(t *testing.T) {
		mockUsers := &MockUsers{}
		mockUsers.On("FindUser", context.Background(), mock.Anything).Return(domain.User{}, errors.New("user not found"))

		app := &App{
			users: mockUsers,
		}

		user, err := app.SignIn(context.Background(), "username", "password")
		assert.ErrorContains(t, err, "failed to find a user: user not found")
		assert.Equal(t, domain.User{}, user)
	})

	t.Run("user found", func(t *testing.T) {
		hashedPassword, err := apputil.HashPassword("password")
		assert.NoError(t, err)

		mockUsers := &MockUsers{}
		mockUsers.On("FindUser", context.Background(), mock.Anything).Return(domain.User{PasswordHash: hashedPassword}, nil)

		app := &App{
			users: mockUsers,
		}

		user, err := app.SignIn(context.Background(), "username", "password")
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
}
