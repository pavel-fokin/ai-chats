package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"ai-chats/internal/domain"
	pkgJson "ai-chats/internal/pkg/json"
	"ai-chats/internal/pkg/types"
)

type Chat struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
}

type Message struct {
	ID        string `json:"id"`
	Sender    string `json:"sender"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
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

type GetOllamaModelsLibraryResponse struct {
	ModelCards []domain.ModelCard `json:"modelCards"`
}

func NewGetOllamaModelsLibraryResponse(modelCards []*domain.ModelCard) GetOllamaModelsLibraryResponse {
	var res GetOllamaModelsLibraryResponse
	for _, modelCard := range modelCards {
		res.ModelCards = append(res.ModelCards, domain.ModelCard{
			ModelName:   modelCard.ModelName,
			Description: modelCard.Description,
			Tags:        modelCard.Tags,
		})
	}
	return res
}

// WriteSuccessResponse writes a JSON response.
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

// WriteErrorResponse writes a JSON response.
func WriteErrorResponse(w http.ResponseWriter, statusCode int, errs ...Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{Errors: errs}

	json.NewEncoder(w).Encode(response)
}

// WriteServerSentEvent writes a server sent event to the response.
func WriteServerSentEvent(w http.ResponseWriter, event types.Message) error {
	if event == nil {
		// Ignore nil events.
		return nil
	}
	fmt.Fprintf(w, "event: %s\n", event.Type())
	fmt.Fprintf(w, "data: %s\n\n", pkgJson.MustMarshal(context.Background(), event))
	return nil
}
