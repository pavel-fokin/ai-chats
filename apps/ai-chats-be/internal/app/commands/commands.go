package commands

import "ai-chats/internal/domain/events"

const (
	GenerateTitleType   events.EventType = "GenerateTitle"
	PullOllamaModelType events.EventType = "PullOllamaModel"
)

type PullOllamaModel struct {
	Model string `json:"model"`
}

func NewPullOllamaModel(model string) PullOllamaModel {
	return PullOllamaModel{
		Model: model,
	}
}

func (PullOllamaModel) Type() events.EventType {
	return PullOllamaModelType
}

type GenerateChatTitle struct {
	ChatID string `json:"chatID"`
}

func NewGenerateChatTitle(chatID string) GenerateChatTitle {
	return GenerateChatTitle{
		ChatID: chatID,
	}
}

func (GenerateChatTitle) Type() events.EventType {
	return GenerateTitleType
}
