package api

import (
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"time"
)

type SignInResponse struct {
	AccessToken string `json:"accessToken"`
}

type SignUpResponse struct {
	AccessToken string `json:"accessToken"`
}
type PostMessagesResponse struct {
	Message
}

type GetChatsResponse struct {
	Chats []Chat `json:"chats"`
}

func NewGetChatsResponse(chats []domain.Chat) GetChatsResponse {
	var res GetChatsResponse
	for _, chat := range chats {
		res.Chats = append(res.Chats, Chat{
			ID:        chat.ID.String(),
			Title:     chat.Title,
			CreatedAt: chat.CreatedAt.Format(time.RFC3339Nano),
		})
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

type GetChatResponse struct {
	Chat Chat `json:"chat"`
}

func NewGetChatResponse(chat domain.Chat) GetChatResponse {
	return GetChatResponse{
		Chat: Chat{
			ID:        chat.ID.String(),
			Title:     chat.Title,
			CreatedAt: chat.CreatedAt.Format(time.RFC3339Nano),
		},
	}
}

type PostChatsResponse struct {
	Chat Chat `json:"chat"`
}

func NewPostChatsResponse(chat domain.Chat) PostChatsResponse {
	return PostChatsResponse{
		Chat: Chat{
			ID:        chat.ID.String(),
			Title:     chat.Title,
			CreatedAt: chat.CreatedAt.Format(time.RFC3339Nano),
		},
	}
}

type GetOllamaModelsResponse struct {
	Models []domain.Model `json:"models"`
}

func NewGetOllamaModelsResponse(models []domain.Model) GetOllamaModelsResponse {
	return GetOllamaModelsResponse{Models: models}
}
