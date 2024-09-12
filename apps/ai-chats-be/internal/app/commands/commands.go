package commands

type PullOllamaModel struct {
	Model string `json:"model"`
}

func NewPullOllamaModel(model string) PullOllamaModel {
	return PullOllamaModel{
		Model: model,
	}
}

type GenerateChatTitle struct {
	ChatID string `json:"chatID"`
}

func NewGenerateChatTitle(chatID string) GenerateChatTitle {
	return GenerateChatTitle{
		ChatID: chatID,
	}
}
