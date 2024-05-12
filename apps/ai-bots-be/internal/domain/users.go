package domain

import "context"

// Users represents a repository of users.
type Users interface {
	AddUser(ctx context.Context, user User) error
	FindUser(ctx context.Context, username string) (User, error)
}
