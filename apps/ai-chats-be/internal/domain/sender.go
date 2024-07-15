package domain

import (
	"fmt"
	"strings"
)

// Sender represents a sender of a message.
type Sender struct {
	sender string
}

// NewUserSender creates a new user sender.
func NewUserSender(user User) Sender {
	return Sender{sender: fmt.Sprintf("user:%s", user.ID)}
}

// NewModelSender creates a new model sender.
func NewModelSender(model Model) Sender {
	return Sender{sender: model.String()}
}

// String returns the string representation of the sender.
func (s Sender) String() string {
	return s.sender
}

// IsUser returns true if the sender is a user.
func (s Sender) IsUser() bool {
	return strings.HasPrefix(s.sender, "user:")
}

// IsModel returns true if the sender is a model.
func (s Sender) IsModel() bool {
	return strings.HasPrefix(s.sender, "model:")
}
