package events

import "EventDrivenArchitectureGoLang/src/main/domain/event"

type EventHandler interface {
	Handle(event event.Event)
}
