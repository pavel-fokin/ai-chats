package domain

import (
	"github.com/google/uuid"

	"ai-chats/internal/pkg/crypto"
)

type UserID = uuid.UUID

type User struct {
	ID           UserID
	Username     string
	PasswordHash string
}

func NewUser(username, password string) User {
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		panic(err)
	}

	return User{
		ID:           uuid.New(),
		Username:     username,
		PasswordHash: hashedPassword,
	}
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return err
	}

	u.PasswordHash = hashedPassword
	return nil
}

// VerifyPassword compares a password with a hash.
func (u User) VerifyPassword(password string) error {
	if err := crypto.VerifyPassword(u.PasswordHash, password); err != nil {
		return ErrInvalidPassword
	}
	return nil
}
