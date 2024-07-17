//go:build !race

package app

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/crypto"
)

type MockUsers struct {
	mock.Mock
}

func (m *MockUsers) Add(ctx context.Context, user domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUsers) FindByUsernameWithPassword(ctx context.Context, username string) (domain.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUsers) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestAppSignUp(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	crypto.InitBcryptCost(1)

	mockUsers := &MockUsers{}
	mockUsers.On("Add", ctx, mock.AnythingOfType("User")).Return(nil)

	app := &App{
		users: mockUsers,
	}

	user, err := app.SignUp(ctx, "username", "password")
	assert.NoError(err)
	assert.NotNil(user)

	// Verify that password is not stored in plain text.
	assert.NotEqual("password", user.PasswordHash)

	t.Run("failed to set up a password", func(t *testing.T) {
		t.Skip("This test is not implemented yet")
		app := &App{
			users: mockUsers,
		}

		user, err := app.SignUp(ctx, "username", "")
		assert.ErrorContains(err, "failed to set up a password")
		assert.Equal(domain.User{}, user)
	})

	t.Run("failed to add a user", func(t *testing.T) {
		mockUsers := &MockUsers{}
		mockUsers.On("Add", ctx, mock.AnythingOfType("User")).
			Return(errors.New("failed to add a user"))

		app := &App{
			users: mockUsers,
		}

		user, err := app.SignUp(ctx, "username", "password")
		assert.ErrorContains(err, "failed to add a user")
		assert.Equal(domain.User{}, user)
	})
}

func TestAppLogIn(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	crypto.InitBcryptCost(1)

	t.Run("user not found", func(t *testing.T) {
		mockUsers := &MockUsers{}
		mockUsers.On("FindByUsernameWithPassword", ctx, "username").
			Return(domain.User{}, domain.ErrUserNotFound)

		app := &App{
			users: mockUsers,
		}

		user, err := app.LogIn(ctx, "username", "password")
		assert.ErrorIs(err, domain.ErrUserNotFound)
		assert.Equal(domain.User{}, user)
	})

	t.Run("success", func(t *testing.T) {
		hashedPassword, err := crypto.HashPassword("password")
		assert.NoError(err)

		mockUsers := &MockUsers{}
		mockUsers.On("FindByUsernameWithPassword", ctx, "username").
			Return(domain.User{PasswordHash: hashedPassword}, nil)

		app := &App{
			users: mockUsers,
		}

		user, err := app.LogIn(ctx, "username", "password")
		assert.NoError(err)
		assert.NotNil(user)
	})

	t.Run("failed to verify password", func(t *testing.T) {
		hashedPassword, err := crypto.HashPassword("password")
		assert.NoError(err)

		mockUsers := &MockUsers{}
		mockUsers.On("FindByUsernameWithPassword", ctx, "username").
			Return(domain.User{PasswordHash: hashedPassword}, nil)

		app := &App{
			users: mockUsers,
		}

		user, err := app.LogIn(ctx, "username", "wrong_password")
		assert.ErrorContains(err, "failed to verify password")
		assert.Equal(domain.User{}, user)
	})
}
