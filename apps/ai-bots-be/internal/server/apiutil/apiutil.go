package apiutil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Error struct {
	Message string `json:"message"`
}

// ErrorResponse.
type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

// SuccessResponse is a wrapper arround payload.
type SuccessResponse struct {
	Data   any     `json:"data,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

// AsErrorResponse.
func AsErrorResponse(
	w http.ResponseWriter, err error, statusCode int,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Make error response.
	payload := ErrorResponse{}
	payload.Errors = []Error{{Message: fmt.Sprint(err)}}

	// Encode json.
	json.NewEncoder(w).Encode(payload)
}

func AsSuccessResponse[T any](
	w http.ResponseWriter, payload T, statusCode int,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	res := SuccessResponse{}
	res.Data = payload

	json.NewEncoder(w).Encode(res)
}

// ParseJSON parses the request body into the given interface.
func ParseJSON(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}
