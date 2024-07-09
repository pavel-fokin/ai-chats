package api

import (
	"errors"
)

var (
	ErrUnauthorized  = errors.New("unauthorized")
	ErrInternal      = errors.New("internal error")
	ErrNotFound      = errors.New("not found")
	ErrBadRequest    = errors.New("bad request")
	ErrUsernameTaken = errors.New("that username is already taken")
)

var (
	NotFound = Error{
		Message: "Not found.",
	}
	BadRequest = Error{
		Message: "Bad request.",
	}
	Unauthorized = Error{
		Message: "Unauthorized.",
	}
	InternalError = Error{
		Message: "Internal error.",
	}
	UsernameIsTaken = Error{
		Message: "That username is already taken. Try another one.",
	}
)
