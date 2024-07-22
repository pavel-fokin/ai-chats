package domain

import (
	"context"
)

// Users represents a repository of users.
type Users interface {
	Add(ctx context.Context, user User) error
	FindByID(ctx context.Context, id UserID) (User, error)
	FindByUsernameWithPassword(ctx context.Context, username string) (User, error)
}
