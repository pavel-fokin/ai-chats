package api

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrInternal     = errors.New("internal error")
)
