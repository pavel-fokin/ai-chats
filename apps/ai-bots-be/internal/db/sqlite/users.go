package sqlite

import (
	"context"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
)

// AddUser adds a new user to the database.
func (db *Sqlite) AddUser(ctx context.Context, user domain.User) error {
	_, err := db.db.ExecContext(
		ctx,
		"INSERT INTO user (id, username, password_hash) VALUES (?, ?, ?);",
		user.ID, user.Username, user.PasswordHash,
	)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

// FindUser finds a user by username.
func (db *Sqlite) FindUser(ctx context.Context, username string) (domain.User, error) {
	user := domain.User{
		Username: username,
	}

	err := db.db.QueryRowContext(
		ctx, "SELECT id, password_hash FROM user WHERE username = ?;", username,
	).Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to select user: %w", err)
	}

	return user, nil
}
