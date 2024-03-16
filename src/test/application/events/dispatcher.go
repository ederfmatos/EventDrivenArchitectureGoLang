package events

type EventDispatcher interface {
	Dispatch(event Event) error
	Register(eventName string, handler EventHandler) error
	UnRegister(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear()
}
