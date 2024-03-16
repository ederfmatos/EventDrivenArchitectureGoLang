package events

import (
	"EventDrivenArchitectureGoLang/src/application/events"
	"errors"
)

var ErrorHandlerAlreadyRegistered = errors.New("event handler already registered")

type DefaultEventDispatcher struct {
	handlers map[string][]events.EventHandler
}

func NewDefaultEventDispatcher() *DefaultEventDispatcher {
	return &DefaultEventDispatcher{handlers: make(map[string][]events.EventHandler)}
}

func (dispatcher *DefaultEventDispatcher) Dispatch(event events.Event) error {
	if dispatcher.handlers[event.GetName()] == nil {
		return nil
	}
	for _, handler := range dispatcher.handlers[event.GetName()] {
		handler.Handle(event)
	}
	return nil
}

func (dispatcher *DefaultEventDispatcher) Register(eventName string, handler events.EventHandler) error {
	if _, ok := dispatcher.handlers[eventName]; ok {
		for _, h := range dispatcher.handlers[eventName] {
			if h == handler {
				return ErrorHandlerAlreadyRegistered
			}
		}
	}
	dispatcher.handlers[eventName] = append(dispatcher.handlers[eventName], handler)
	return nil
}

func (dispatcher *DefaultEventDispatcher) UnRegister(eventName string, handler events.EventHandler) error {
	if dispatcher.handlers[eventName] == nil {
		return nil
	}
	for i, h := range dispatcher.handlers[eventName] {
		if h == handler {
			dispatcher.handlers[eventName] = append(dispatcher.handlers[eventName][:i], dispatcher.handlers[eventName][i+1:]...)
			return nil
		}
	}
	return nil
}

func (dispatcher *DefaultEventDispatcher) Has(eventName string, handler events.EventHandler) bool {
	if dispatcher.handlers[eventName] == nil {
		return false
	}
	for _, h := range dispatcher.handlers[eventName] {
		if h == handler {
			return true
		}
	}
	return false
}

func (dispatcher *DefaultEventDispatcher) Clear() {
	dispatcher.handlers = make(map[string][]events.EventHandler)
}
