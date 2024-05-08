package app

import (
	"context"
	"errors"
	"pavel-fokin/ai/apps/ai-bots-be/internal/app/apputil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAuthDB struct {
	mock.Mock
}

func (m *mockAuthDB) CreateUser(ctx context.Context, username, password string) (User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(User), args.Error(1)
}

func (m *mockAuthDB) FindUser(ctx context.Context, username string) (User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(User), args.Error(1)
}

func TestSignUp(t *testing.T) {
	db := &mockAuthDB{}
	db.On("CreateUser", context.Background(), "username", mock.Anything).Return(User{}, nil)

	app := &App{
		userDB: db,
	}

	user, err := app.SignUp(context.Background(), "username", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	// Verify that password is not stored in plain text.
	assert.NotEqual(t, "password", db.Calls[0].Arguments[2])
}

func TestSignIn(t *testing.T) {
	t.Run("user not found", func(t *testing.T) {
		db := &mockAuthDB{}
		db.On("FindUser", context.Background(), mock.Anything).Return(User{}, errors.New("user not found"))

		app := &App{
			userDB: db,
		}

		user, err := app.SignIn(context.Background(), "username", "password")
		assert.ErrorContains(t, err, "failed to sign in a user: user not found")
		assert.Equal(t, User{}, user)
	})

	t.Run("user found", func(t *testing.T) {
		hashedPassword, err := apputil.HashPassword("password")
		assert.NoError(t, err)

		db := &mockAuthDB{}
		db.On("FindUser", context.Background(), mock.Anything).Return(User{Password: hashedPassword}, nil)

		app := &App{
			userDB: db,
		}

		user, err := app.SignIn(context.Background(), "username", "password")
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
}
