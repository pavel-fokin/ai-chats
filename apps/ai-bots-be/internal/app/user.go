package app

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app/apputil"

	"github.com/google/uuid"
)

// User is an application model that represents a user.
type User struct {
	ID       uuid.UUID
	Username string
	Password string
}

// SignIn signs in a user.
func (a *App) SignIn(ctx context.Context, username, password string) (User, error) {
	user, err := a.userDB.FindUser(ctx, username)
	if err != nil {
		return User{}, fmt.Errorf("failed to sign in a user: %w", err)
	}

	if err := apputil.VerifyPassword(user.Password, password); err != nil {
		return User{}, fmt.Errorf("failed to sign in a user: %w", err)
	}

	return user, nil
}

// SignUp signs up a user.
func (a *App) SignUp(ctx context.Context, username, password string) (User, error) {
	hashedPassword, err := apputil.HashPassword(password)
	if err != nil {
		return User{}, fmt.Errorf("failed to sign up a user: %w", err)
	}

	user, err := a.userDB.CreateUser(ctx, username, hashedPassword)
	if err != nil {
		return User{}, fmt.Errorf("failed to sign up a user: %w", err)
	}

	return user, nil
}
