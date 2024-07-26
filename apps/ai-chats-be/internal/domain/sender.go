package domain

import (
	"fmt"
	"strings"
)

// Sender represents a sender of a message.
type Sender struct {
	sender string
}

func NewSender(sender string) Sender {
	return Sender{sender: sender}
}

// NewUserSender creates a new user sender.
func NewUserSender(userID UserID) Sender {
	return NewSender(fmt.Sprintf("user:%s", userID))
}

// NewModelSender creates a new model sender.
func NewModelSender(modelID ModelID) Sender {
	return NewSender(fmt.Sprintf("model:%s", modelID))
}

// String returns the string representation of the sender.
func (s Sender) String() string {
	return s.sender
}

// MarshalJSON returns the JSON representation of the sender.
func (s Sender) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, s.sender)), nil
}

// UnmarshalJSON parses the JSON representation of the sender.
func (s *Sender) UnmarshalJSON(data []byte) error {
	s.sender = strings.Trim(string(data), `"`)
	return nil
}

// Format returns the sender in the format "type:id".
func (s Sender) Format() string {
	if s.IsUser() {
		return "You"
	} else if s.IsModel() {
		parts := strings.Split(s.sender, ":")
		if len(parts) == 3 {
			return fmt.Sprintf("Model (%s:%s)", parts[1], parts[2])
		} else if len(parts) == 2 {
			return fmt.Sprintf("Model (%s)", parts[1])
		} else {
			return "Model"
		}
	}
	return ""
}

// IsUser returns true if the sender is a user.
func (s Sender) IsUser() bool {
	return strings.HasPrefix(s.sender, "user:")
}

// IsModel returns true if the sender is a model.
func (s Sender) IsModel() bool {
	return strings.HasPrefix(s.sender, "model:")
}
