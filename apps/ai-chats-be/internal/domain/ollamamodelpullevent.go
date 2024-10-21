package domain

import (
	"time"

	"ai-chats/internal/pkg/types"

	"github.com/google/uuid"
)

type OllamaPullingFinalStatus string

const (
	OllamaModelPullStartedType   types.MessageType = "OllamaModelPullStarted"
	OllamaModelPullCompletedType types.MessageType = "OllamaModelPullCompleted"
	OllamaModelPullFailedType    types.MessageType = "OllamaModelPullFailed"
)

type OllamaModelPullEvent interface {
	ID() uuid.UUID
	Type() types.MessageType
	OccurredAt() time.Time
	Model() string
}

type OllamaModelPullStarted struct {
	id         uuid.UUID
	model      string
	occurredAt time.Time
}

func NewOllamaModelPullStarted(model string) OllamaModelPullStarted {
	return OllamaModelPullStarted{
		id:         uuid.New(),
		model:      model,
		occurredAt: time.Now().UTC(),
	}
}

func (e OllamaModelPullStarted) ID() uuid.UUID {
	return e.id
}

func (e OllamaModelPullStarted) OccurredAt() time.Time {
	return e.occurredAt
}

func (e OllamaModelPullStarted) Type() types.MessageType {
	return OllamaModelPullStartedType
}

func (e OllamaModelPullStarted) Model() string {
	return e.model
}

type OllamaModelPullCompleted struct {
	id         uuid.UUID
	model      string
	occurredAt time.Time
}

func NewOllamaModelPullCompleted(model string) OllamaModelPullCompleted {
	return OllamaModelPullCompleted{
		id:         uuid.New(),
		model:      model,
		occurredAt: time.Now().UTC(),
	}
}

func (e OllamaModelPullCompleted) ID() uuid.UUID {
	return e.id
}

func (e OllamaModelPullCompleted) OccurredAt() time.Time {
	return e.occurredAt
}

func (e OllamaModelPullCompleted) Type() types.MessageType {
	return OllamaModelPullCompletedType
}

func (e OllamaModelPullCompleted) Model() string {
	return e.model
}

type OllamaModelPullFailed struct {
	id         uuid.UUID
	model      string
	occurredAt time.Time
}

func NewOllamaModelPullFailed(model string) OllamaModelPullFailed {
	return OllamaModelPullFailed{
		id:         uuid.New(),
		model:      model,
		occurredAt: time.Now().UTC(),
	}
}

func (e OllamaModelPullFailed) ID() uuid.UUID {
	return e.id
}

func (e OllamaModelPullFailed) Type() types.MessageType {
	return OllamaModelPullFailedType
}

func (e OllamaModelPullFailed) OccurredAt() time.Time {
	return e.occurredAt
}

func (e OllamaModelPullFailed) Model() string {
	return e.model
}
