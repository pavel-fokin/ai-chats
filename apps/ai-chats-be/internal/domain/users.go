package domain

import (
	"context"

	"github.com/google/uuid"
)

// Users represents a repository of users.
type Users interface {
	Add(ctx context.Context, user User) error
	FindByID(ctx context.Context, id uuid.UUID) (User, error)
	FindByUsernameWithPassword(ctx context.Context, username string) (User, error)
}
