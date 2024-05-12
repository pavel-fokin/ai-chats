package app

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

// SignIn signs in a user.
func (a *App) SignIn(ctx context.Context, username, password string) (domain.User, error) {
	user, err := a.users.FindUser(ctx, username)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to sign in a user: %w", err)
	}

	if err := user.VerifyPassword(password); err != nil {
		return domain.User{}, fmt.Errorf("failed to sign in a user: %w", err)
	}

	return user, nil
}

// SignUp signs up a user.
func (a *App) SignUp(ctx context.Context, username, password string) (domain.User, error) {
	user := domain.NewUser(username)

	pUser := &user
	if err := pUser.SetPassword(password); err != nil {
		return domain.User{}, fmt.Errorf("failed to sign up a user: %w", err)
	}

	if err := a.users.AddUser(ctx, user); err != nil {
		return domain.User{}, fmt.Errorf("failed to sign up a user: %w", err)
	}

	return user, nil
}
