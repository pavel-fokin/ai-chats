package sqlite

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app"

	"github.com/google/uuid"
)

// CreateUser creates a new user.
func (db *Sqlite) CreateUser(ctx context.Context, username, password string) (app.User, error) {
	userID := uuid.New()

	if err := db.db.QueryRowContext(
		ctx,
		"INSERT INTO user (id, username, password) VALUES (?, ?, ?) RETURNING password;",
		userID, username, password,
	).Scan(&password); err != nil {
		return app.User{}, fmt.Errorf("failed to insert user: %w", err)
	}

	return app.User{
		ID:       userID,
		Username: username,
		Password: password,
	}, nil
}

// FindUser finds a user by username.
func (db *Sqlite) FindUser(ctx context.Context, username string) (app.User, error) {
	user := app.User{
		Username: username,
	}

	err := db.db.QueryRowContext(
		ctx, "SELECT id, password FROM user WHERE username = ?;", username,
	).Scan(&user.ID, &user.Password)
	if err != nil {
		return app.User{}, fmt.Errorf("failed to select user: %w", err)
	}

	return user, nil
}
