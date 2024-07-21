package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ai-chats/internal/domain"
	"time"
)

type Chat struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}

type Message struct {
	ID        string `json:"id"`
	Sender    string `json:"sender"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type Error struct {
	Field   string `json:"field,omitempty"`
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

type Response struct {
	Data   any     `json:"data,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

func NewResponse(data any, errors []Error) Response {
	return Response{Data: data, Errors: errors}
}

type LogInResponse struct {
	AccessToken string `json:"accessToken"`
}

type SignUpResponse struct {
	AccessToken string `json:"accessToken"`
}

func NewSignUpResponse(accessToken string) SignUpResponse {
	return SignUpResponse{AccessToken: accessToken}
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
		res.Messages = append(res.Messages, Message{
			ID:        message.ID.String(),
			Text:      message.Text,
			Sender:    message.Sender.Format(),
			CreatedAt: message.CreatedAt.Format(time.RFC3339Nano),
		})
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
	Models []domain.OllamaModel `json:"models"`
}

func NewGetOllamaModelsResponse(models []domain.OllamaModel) GetOllamaModelsResponse {
	return GetOllamaModelsResponse{Models: models}
}

// AsErrorResponse.
func AsErrorResponse(
	w http.ResponseWriter, err error, statusCode int,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	res := Response{
		Errors: []Error{{Message: err.Error()}},
	}

	json.NewEncoder(w).Encode(res)
}

func AsSuccessResponse(
	w http.ResponseWriter, data any, statusCode int,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		if statusCode != http.StatusNoContent {
			panic("payload is nil")
		}
		return
	}

	res := Response{Data: data}

	json.NewEncoder(w).Encode(res)
}

// WriteResponse writes a JSON response.
func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		if statusCode != http.StatusNoContent {
			panic("payload is nil")
		}
		return
	}

	response := Response{Data: data}

	json.NewEncoder(w).Encode(response)
}

// WriteResponse writes a JSON response.
func WriteErrorResponse(w http.ResponseWriter, statusCode int, errs ...Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{Errors: errs}

	json.NewEncoder(w).Encode(response)
}

// WriteEvent writes a server sent event to the response.
func WriteEvent(w http.ResponseWriter, data []byte) error {
	fmt.Fprintf(w, "data: %s\n\n", data)
	return nil
}
