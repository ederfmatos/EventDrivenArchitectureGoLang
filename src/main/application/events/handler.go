package events

type EventHandler interface {
	EventName() string
	Handle(message []byte) error
}
