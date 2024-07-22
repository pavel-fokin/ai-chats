package domain

import (
	"github.com/google/uuid"

	"ai-chats/internal/pkg/crypto"
)

type UserID = uuid.UUID

func NewUserID() UserID {
	return uuid.New()
}

type User struct {
	ID           UserID
	Username     string
	PasswordHash string
}

func NewUser(username string) User {
	return User{
		ID:       NewUserID(),
		Username: username,
	}
}

func NewUserWithPassword(username, password string, hashCost int) User {
	hashedPassword, err := crypto.HashPassword(password, hashCost)
	if err != nil {
		panic(err)
	}

	return User{
		ID:           NewUserID(),
		Username:     username,
		PasswordHash: hashedPassword,
	}
}

// VerifyPassword compares a password with a hash.
func (u User) VerifyPassword(password string) error {
	if err := crypto.VerifyPassword(u.PasswordHash, password); err != nil {
		return ErrInvalidPassword
	}
	return nil
}
