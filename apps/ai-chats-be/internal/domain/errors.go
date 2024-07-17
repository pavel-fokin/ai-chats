package domain

import "errors"

var (
	ErrChatNotFound      = errors.New("chat not found")
	ErrTagNotFound       = errors.New("tag not found")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserAlreadyExists = errors.New("user already exists")
)
