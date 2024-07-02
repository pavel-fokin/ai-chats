package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pavel-fokin/ai/apps/ai-bots-be/internal/domain"
	"time"
)

type Error struct {
	Message string `json:"message"`
}

// ErrorResponse.
type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

// SuccessResponse is a wrapper arround payload.
type SuccessResponse struct {
	Data   any     `json:"data,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

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

// AsErrorResponse.
func AsErrorResponse(
	w http.ResponseWriter, err error, statusCode int,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Make error response.
	payload := ErrorResponse{}
	payload.Errors = []Error{{Message: fmt.Sprint(err)}}

	// Encode json.
	json.NewEncoder(w).Encode(payload)
}

func AsSuccessResponse(
	w http.ResponseWriter, payload any, statusCode int,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if payload == nil {
		if statusCode != http.StatusNoContent {
			panic("payload is nil")
		}
		return
	}

	res := SuccessResponse{Data: payload}

	json.NewEncoder(w).Encode(res)
}

// WriteEvent writes an server sent event to the response.
func WriteEvent(w http.ResponseWriter, data []byte) error {
	// Encode json.
	// dataJSON, err := json.Marshal(data)
	// if err != nil {
	// 	return fmt.Errorf("failed to encode event: %w", err)
	// }

	fmt.Fprintf(w, "data: %s\n\n", data)
	return nil
}
