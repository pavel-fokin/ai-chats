package commands

type PullOllamaModel struct {
	Model string `json:"model"`
}

func NewPullOllamaModel(model string) PullOllamaModel {
	return PullOllamaModel{
		Model: model,
	}
}
