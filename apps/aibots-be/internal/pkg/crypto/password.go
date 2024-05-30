package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

var bcryptCost int

func InitBcryptCost(value int) {
	bcryptCost = value
}

// HashPassword hashes a password.
func HashPassword(password string) (string, error) {
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassowrd), nil
}

// VerifyPassword compares a password with a hash.
func VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
