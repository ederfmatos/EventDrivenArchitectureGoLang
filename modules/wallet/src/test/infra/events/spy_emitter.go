package events

import "wallet/src/main/domain/event"

type SpyEventEmitter struct {
	Calls int
}

func NewSpyEventEmitter() *SpyEventEmitter {
	return &SpyEventEmitter{Calls: 0}
}

func (emitter *SpyEventEmitter) Emit(event event.Event) error {
	emitter.Calls++
	return nil
}
