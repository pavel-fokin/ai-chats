package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 14

type User struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
}

func NewUser(username string) User {
	return User{
		ID:       uuid.New(),
		Username: username,
	}
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	u.PasswordHash = hashedPassword
	return nil
}

// VerifyPassword compares a password with a hash.
func (u User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

// HashPassword hashes a password.
func hashPassword(password string) (string, error) {
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassowrd), nil
}
