package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"

	"github.com/google/uuid"
)

// Users represents a repository of users.
type Users struct {
	DB
}

// NewUsers creates a new users repository.
func NewUsers(db *sql.DB) *Users {
	return &Users{DB{db: db}}
}

// AddUser adds a new user to the database.
func (u *Users) Add(ctx context.Context, user domain.User) error {
	_, err := u.DBTX(ctx).ExecContext(
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
func (u *Users) FindByUsernameWithPassword(ctx context.Context, username string) (domain.User, error) {
	user := domain.User{
		Username: username,
	}

	err := u.DBTX(ctx).QueryRowContext(
		ctx, "SELECT id, password_hash FROM user WHERE username = ?;", username,
	).Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to select user: %w", err)
	}

	return user, nil
}

// FindByID finds a user by ID.
func (u *Users) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	user := domain.User{
		ID: id,
	}

	err := u.DBTX(ctx).QueryRowContext(
		ctx, "SELECT username FROM user WHERE id = ?;", id,
	).Scan(&user.Username)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to find user by id: %w", err)
	}

	return user, nil
}
