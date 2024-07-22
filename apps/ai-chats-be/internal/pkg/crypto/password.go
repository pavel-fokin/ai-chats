package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password.
func HashPassword(password string, hashCost int) (string, error) {
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassowrd), nil
}

// VerifyPassword compares a password with a hash.
func VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
