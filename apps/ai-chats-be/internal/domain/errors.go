package domain

import "errors"

var (
	ErrChatNotFound        = errors.New("chat not found")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrOllamaModelNotFound = errors.New("ollama model not found")
	ErrTagNotFound         = errors.New("tag not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrUserNotFound        = errors.New("user not found")
)
