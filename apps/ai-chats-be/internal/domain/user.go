package domain

import (
	"github.com/google/uuid"

	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/crypto"
)

type User struct {
	ID           uuid.UUID
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
	return crypto.VerifyPassword(u.PasswordHash, password)
}
