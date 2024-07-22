package domain

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"

	"ai-chats/internal/pkg/crypto"
)

// UserID represents a unique identifier for a user.
type UserID uuid.UUID

func NewUserID() UserID {
	return UserID(uuid.New())
}

func (id UserID) String() string {
	return uuid.UUID(id).String()
}

func (id UserID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(id))
}

func (id *UserID) UnmarshalJSON(data []byte) error {
	var u uuid.UUID
	if err := json.Unmarshal(data, &u); err != nil {
		return err
	}
	*id = UserID(u)
	return nil
}

func (id *UserID) Scan(value interface{}) error {
	u, err := uuid.Parse(value.(string))
	if err != nil {
		return err
	}
	*id = UserID(u)
	return nil
}

func (id UserID) Value() (driver.Value, error) {
	return uuid.UUID(id).String(), nil
}

// User represents a user in the system.
type User struct {
	ID           UserID
	Username     string
	PasswordHash string
}

// NewUser creates a new User with the given username.
func NewUser(username string) User {
	return User{
		ID:       NewUserID(),
		Username: username,
	}
}

// NewUserWithPassword creates a new User with the provided username, password, and hash cost.
// It hashes the password using the specified hash cost and returns the created User.
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
