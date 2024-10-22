package domain

import "errors"

var (
	ErrChatNotFound      = errors.New("chat not found")
	ErrChatAccessDenied  = errors.New("chat access denied")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)
