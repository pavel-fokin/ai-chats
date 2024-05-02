package api

type SignInResponse struct {
	AccessToken string `json:"accessToken"`
}

type SignUpResponse struct {
	AccessToken string `json:"accessToken"`
}

type PostMessagesRequest struct {
	Message
}
