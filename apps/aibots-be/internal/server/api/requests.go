package api

type PostMessagesRequest struct {
	Message
}

type SignInRequest struct {
	UserCredentials
}

type SignUpRequest struct {
	UserCredentials
}
