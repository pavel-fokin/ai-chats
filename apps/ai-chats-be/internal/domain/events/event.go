package events

type Event interface {
	Type() EventType
}
