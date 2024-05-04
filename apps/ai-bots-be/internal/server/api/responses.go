package api

import "pavel-fokin/ai/apps/ai-bots-be/internal/app/domain"

type UserCredentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInRequest struct {
	UserCredentials
}

type SignUpRequest struct {
	UserCredentials
}

type PostMessagesResponse struct {
	Message
}

type Chat struct {
	ID string `json:"id"`
}

type GetChatsResponse struct {
	Chats []Chat `json:"chats"`
}

func NewGetChatsResponse(chats []domain.Chat) GetChatsResponse {
	var res GetChatsResponse
	for _, chat := range chats {
		res.Chats = append(res.Chats, Chat{ID: chat.ID.String()})
	}
	return res
}

type GetMessagesResponse struct {
	Messages []Message `json:"messages"`
}

func NewGetMessagesResponse(messages []domain.Message) GetMessagesResponse {
	var res GetMessagesResponse
	for _, message := range messages {
		res.Messages = append(res.Messages, Message{ID: message.ID.String(), Text: message.Text})
	}
	return res
}