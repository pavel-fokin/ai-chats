package commands

import (
	"ai-chats/internal/pkg/types"
)

const (
	GenerateTitleType   types.MessageType = "GenerateTitle"
	PullOllamaModelType types.MessageType = "PullOllamaModel"
)

type PullOllamaModel struct {
	Model string `json:"model"`
}

func NewPullOllamaModel(model string) PullOllamaModel {
	return PullOllamaModel{
		Model: model,
	}
}

func (PullOllamaModel) Type() types.MessageType {
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

func (GenerateChatTitle) Type() types.MessageType {
	return GenerateTitleType
}
