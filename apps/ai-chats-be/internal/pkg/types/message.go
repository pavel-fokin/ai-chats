package types

type MessageType string

// Message is an abstract message interface for all type of messages in the app.
type Message interface {
	Type() MessageType
}
