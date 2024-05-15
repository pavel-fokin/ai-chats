package worker

import "github.com/google/uuid"

// GenerateResponse represents a request to generate a LLM response.
type GenerateResponse struct {
	ChatID uuid.UUID `json:"chat_id"`
}
