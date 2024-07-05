package app

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

// LogIn logs in a user
func (a *App) LogIn(ctx context.Context, username, password string) (domain.User, error) {
	user, err := a.users.FindByUsernameWithPassword(ctx, username)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to find a user: %w", err)
	}

	if err := user.VerifyPassword(password); err != nil {
		return domain.User{}, fmt.Errorf("failed to verify password: %w", err)
	}

	return user, nil
}

// SignUp signs up a user.
func (a *App) SignUp(ctx context.Context, username, password string) (domain.User, error) {
	user := domain.NewUser(username)

	pUser := &user
	if err := pUser.SetPassword(password); err != nil {
		return domain.User{}, fmt.Errorf("failed to set up a password: %w", err)
	}

	if err := a.users.Add(ctx, user); err != nil {
		switch err {
		case domain.ErrUserAlreadyExists:
			return domain.User{}, fmt.Errorf("failed to add a user: %w", domain.ErrUserAlreadyExists)
		default:
			return domain.User{}, fmt.Errorf("failed to add a user: %w", err)
		}
	}

	return user, nil
}
