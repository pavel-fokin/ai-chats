package api

import (
	"errors"
)

var (
	ErrInternal   = errors.New("internal error")
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
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
	UsernameOrPasswordIsIncorrect = Error{
		Message: "Username or password is incorrect. Please try again.",
	}
)
