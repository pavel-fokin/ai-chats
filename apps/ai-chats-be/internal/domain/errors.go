package domain

import "errors"

var (
	ErrChatNotFound = errors.New("chat not found")
	ErrTagNotFound  = errors.New("tag not found")
)