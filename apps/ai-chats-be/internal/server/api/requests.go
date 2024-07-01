package api

type PostChatsRequest struct {
	Message string `json:"message"`
}

type PostMessagesRequest struct {
	Message
}

type SignInRequest struct {
	UserCredentials
}

type SignUpRequest struct {
	UserCredentials
}

type PostOllamaModelsRequest struct {
	Model string `json:"model"`
}
