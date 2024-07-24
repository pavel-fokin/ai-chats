package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"ai-chats/internal/domain"
)

// Users represents a repository of users.
type Users struct {
	DB
}

// NewUsers creates a new users repository.
func NewUsers(db *sql.DB) *Users {
	return &Users{DB{db: db}}
}

// Add adds a new user to the database.
func (u *Users) Add(ctx context.Context, user domain.User) error {
	_, err := u.DBTX(ctx).ExecContext(
		ctx,
		"INSERT INTO user (id, username, password_hash) VALUES (?, ?, ?);",
		user.ID, user.Username, user.PasswordHash,
	)
	if err != nil {
		switch {
		case isUniqueConstraintViolation(err):
			return domain.ErrUserAlreadyExists
		default:
			return fmt.Errorf("failed to insert user: %w", err)
		}
	}

	return nil
}

// FindByUsernameWithPassword retrieves a user from the database by their username,
// including their password hash. It returns the found user and any error encountered.
// If the user is not found, it returns domain.ErrUserNotFound.
func (u *Users) FindByUsernameWithPassword(ctx context.Context, username string) (domain.User, error) {
	user := domain.User{
		Username: username,
	}

	err := u.DBTX(ctx).QueryRowContext(
		ctx, "SELECT id, password_hash FROM user WHERE username = ?;", username,
	).Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return domain.User{}, domain.ErrUserNotFound
		default:
			return domain.User{}, fmt.Errorf("failed to find user by username: %w", err)
		}
	}

	return user, nil
}

// FindByID finds a user by ID.
func (u *Users) FindByID(ctx context.Context, id domain.UserID) (domain.User, error) {
	user := domain.User{
		ID: id,
	}

	err := u.DBTX(ctx).QueryRowContext(
		ctx, "SELECT username FROM user WHERE id = ?;", id,
	).Scan(&user.Username)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return domain.User{}, domain.ErrUserNotFound
		default:
			return domain.User{}, fmt.Errorf("failed to find user by id: %w", err)
		}
	}

	return user, nil
}
