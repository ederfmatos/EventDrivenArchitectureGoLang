package events

import (
	"EventDrivenArchitectureGoLang/src/main/application/events"
	"EventDrivenArchitectureGoLang/src/main/domain/event"
	"errors"
)

var ErrorHandlerAlreadyRegistered = errors.New("event handler already registered")

type DefaultEventDispatcher struct {
	Handlers map[string][]events.EventHandler
}

func NewDefaultEventDispatcher() *DefaultEventDispatcher {
	return &DefaultEventDispatcher{Handlers: make(map[string][]events.EventHandler)}
}

func (dispatcher *DefaultEventDispatcher) Dispatch(event event.Event) error {
	if dispatcher.Handlers[event.GetName()] == nil {
		return nil
	}
	for _, handler := range dispatcher.Handlers[event.GetName()] {
		handler.Handle(event)
	}
	return nil
}

func (dispatcher *DefaultEventDispatcher) Register(eventName string, handler events.EventHandler) error {
	if _, ok := dispatcher.Handlers[eventName]; ok {
		for _, h := range dispatcher.Handlers[eventName] {
			if h == handler {
				return ErrorHandlerAlreadyRegistered
			}
		}
	}
	dispatcher.Handlers[eventName] = append(dispatcher.Handlers[eventName], handler)
	return nil
}

func (dispatcher *DefaultEventDispatcher) UnRegister(eventName string, handler events.EventHandler) error {
	if dispatcher.Handlers[eventName] == nil {
		return nil
	}
	for i, h := range dispatcher.Handlers[eventName] {
		if h == handler {
			dispatcher.Handlers[eventName] = append(dispatcher.Handlers[eventName][:i], dispatcher.Handlers[eventName][i+1:]...)
			return nil
		}
	}
	return nil
}

func (dispatcher *DefaultEventDispatcher) Has(eventName string, handler events.EventHandler) bool {
	if dispatcher.Handlers[eventName] == nil {
		return false
	}
	for _, h := range dispatcher.Handlers[eventName] {
		if h == handler {
			return true
		}
	}
	return false
}

func (dispatcher *DefaultEventDispatcher) Clear() {
	dispatcher.Handlers = make(map[string][]events.EventHandler)
}
