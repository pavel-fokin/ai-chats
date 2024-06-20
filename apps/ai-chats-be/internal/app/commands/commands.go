package commands

import "github.com/google/uuid"

type GenerateResponse struct {
	ID     uuid.UUID `json:"id"`
	ChatID uuid.UUID `json:"chat_id"`
}

func NewGenerateResponse(chatID uuid.UUID) *GenerateResponse {
	return &GenerateResponse{
		ID:     uuid.New(),
		ChatID: chatID,
	}
}

type GenerateTitle struct {
	ID     uuid.UUID `json:"id"`
	ChatID uuid.UUID `json:"chat_id"`
}

func NewGenerateTitle(chatID uuid.UUID) *GenerateTitle {
	return &GenerateTitle{
		ID:     uuid.New(),
		ChatID: chatID,
	}
}
