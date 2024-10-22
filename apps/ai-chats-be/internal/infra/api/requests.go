package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type PostChatsRequest struct {
	DefaultModel string `json:"defaultModel" validate:"required"`
	Message      string `json:"message" validate:"required"`
}

type PostMessagesRequest struct {
	Text string `json:"text" validate:"required"`
}

type UserCredentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LogInRequest struct {
	UserCredentials
}

type SignUpRequest struct {
	UserCredentials
}

type PostOllamaModelsRequest struct {
	Model string `json:"model"`
}

// ParseJSON validates and parses the request body into the given interface.
func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}
