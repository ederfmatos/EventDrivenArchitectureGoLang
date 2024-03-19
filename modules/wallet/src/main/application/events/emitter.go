package events

import "wallet/src/main/domain/event"

type EventEmitter interface {
	Emit(event event.Event) error
}
