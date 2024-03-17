package events

import "EventDrivenArchitectureGoLang/src/main/domain/event"

type EventEmitter interface {
	Emit(event event.Event) error
}
