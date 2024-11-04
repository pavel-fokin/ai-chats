package app

import (
	"context"
	"fmt"

	"ai-chats/internal/domain"
)

// AuthConfig is the configuration for the authentication service.
type AuthConfig struct {
	HashCost int
}

// Auth is an authentication service.
type Auth struct {
	config AuthConfig
	users  domain.Users
}

// NewAuth creates a new authentication service.
func NewAuth(config AuthConfig, users domain.Users) *Auth {
	return &Auth{users: users, config: config}
}

// LogIn logs in a user
func (a *Auth) LogIn(ctx context.Context, username, password string) (domain.User, error) {
	user, err := a.users.FindByUsernameWithPassword(ctx, username)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to find a user by username: %w", err)
	}

	if err := user.VerifyPassword(password); err != nil {
		return domain.User{}, fmt.Errorf("failed to verify password: %w", err)
	}

	return user, nil
}

// SignUp signs up a user.
func (a *Auth) SignUp(ctx context.Context, username, password string) (domain.User, error) {
	user := domain.NewUserWithPassword(username, password, a.config.HashCost)

	if err := a.users.Add(ctx, user); err != nil {
		return domain.User{}, fmt.Errorf("failed to add a user: %w", err)
	}

	return user, nil
}
