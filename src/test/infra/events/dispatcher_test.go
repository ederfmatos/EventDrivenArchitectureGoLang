package events

import (
	"EventDrivenArchitectureGoLang/src/main/domain/event"
	events2 "EventDrivenArchitectureGoLang/src/main/infra/events"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (event TestEvent) GetName() string {
	return event.Name
}

func (event TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (event TestEvent) GetPayload() interface{} {
	return event.Payload
}

type TestEventHandler struct {
	ID         int
	CallsCount int
}

func NewTestEventHandler(id int) *TestEventHandler {
	return &TestEventHandler{ID: id, CallsCount: 0}
}

func (handler *TestEventHandler) Handle(event.Event) {
	handler.CallsCount++
}

func MakeSut() (*events2.DefaultEventDispatcher, *TestEventHandler, TestEvent) {
	return events2.NewDefaultEventDispatcher(), NewTestEventHandler(1), TestEvent{Name: "EventName", Payload: "EventPayload"}
}

func Test_Should_Register_An_Event_Dispatcher(t *testing.T) {
	dispatcher, handler, event := MakeSut()
	handler2 := NewTestEventHandler(2)

	err := dispatcher.Register(event.Name, handler)
	assert.NoError(t, err, "Error registering event handler")

	quantityOfHandlers := len(dispatcher.Handlers[event.GetName()])
	assert.Equal(t, 1, quantityOfHandlers, "Expected dispatcher.handlers[event.GetName()] = 1")

	err = dispatcher.Register(event.Name, handler)
	assert.ErrorIs(t, err, events2.ErrorHandlerAlreadyRegistered, "Error registering event handler")
	assert.True(t, dispatcher.Has(event.Name, handler), "Expected dispatcher.Has(event.Name, handler) to be true")

	err = dispatcher.Register(event.Name, handler2)
	assert.NoError(t, err, "Error registering event handler")

	quantityOfHandlers = len(dispatcher.Handlers[event.GetName()])
	assert.Equal(t, 2, quantityOfHandlers, "Expected dispatcher.handlers[event.GetName()] = 2")

	err = dispatcher.Register(event.Name, handler2)
	assert.ErrorIs(t, err, events2.ErrorHandlerAlreadyRegistered, "Error registering event handler")
	assert.True(t, dispatcher.Has(event.Name, handler2), "Expected dispatcher.Has(event.Name, handler) to be true")
}

func Test_Should_UnRegister_An_Event_Dispatcher(t *testing.T) {
	dispatcher, handler, event := MakeSut()

	err := dispatcher.Register(event.Name, handler)
	assert.NoError(t, err, "Error registering event handler")

	quantityOfHandlers := len(dispatcher.Handlers[event.GetName()])
	assert.Equal(t, 1, quantityOfHandlers, "Expected dispatcher.handlers[event.GetName()] = 1")
	assert.True(t, dispatcher.Has(event.Name, handler), "Expected dispatcher.Has(event.Name, handler) to be true")

	err = dispatcher.UnRegister(event.Name, handler)
	assert.NoError(t, err, "Error unregistering event handler")
	quantityOfHandlers = len(dispatcher.Handlers[event.GetName()])
	assert.Equal(t, 0, quantityOfHandlers, "Expected dispatcher.handlers[event.GetName()] = 0")
	assert.False(t, dispatcher.Has(event.Name, handler), "Expected dispatcher.Has(event.Name, handler) to be false")
}

func Test_Should_Clear(t *testing.T) {
	dispatcher, handler, event := MakeSut()
	handler2 := NewTestEventHandler(2)

	err := dispatcher.Register(event.Name, handler)
	assert.NoError(t, err, "Error registering event handler")
	err = dispatcher.Register(event.Name, handler2)
	assert.NoError(t, err, "Error registering event handler")

	quantityOfHandlers := len(dispatcher.Handlers[event.GetName()])
	assert.Equal(t, 2, quantityOfHandlers, "Expected dispatcher.handlers[event.GetName()] = 2")

	dispatcher.Clear()

	quantityOfHandlers = len(dispatcher.Handlers[event.GetName()])
	assert.Equal(t, 0, quantityOfHandlers, "Expected dispatcher.handlers[event.GetName()] = 0")
	assert.False(t, dispatcher.Has(event.Name, handler), "Expected dispatcher.Has(event.Name, handler) to be false")
	assert.False(t, dispatcher.Has(event.Name, handler2), "Expected dispatcher.Has(event.Name, handler) to be false")
}

func Test_DefaultEventDispatcher_Has(t *testing.T) {
	dispatcher, handler, event := MakeSut()
	assert.False(t, dispatcher.Has(event.Name, handler), "Expected dispatcher.Has(event.Name) to return false")
	err := dispatcher.Register(event.Name, handler)
	assert.NoError(t, err, "Error registering event handler")
	assert.True(t, dispatcher.Has(event.Name, handler), "Expected dispatcher.Has(event.Name) to return true")
}

func Test_DefaultEventDispatcher_Dispatch(t *testing.T) {
	dispatcher, handler, event := MakeSut()
	handler2 := NewTestEventHandler(2)

	err := dispatcher.Register(event.Name, handler)
	assert.NoError(t, err, "Error registering event handler")
	err = dispatcher.Register(event.Name, handler2)
	assert.NoError(t, err, "Error registering event handler")

	err = dispatcher.Dispatch(event)
	assert.NoError(t, err, "Error when dispatch event")

	assert.Equal(t, 1, handler.CallsCount)
	assert.Equal(t, 1, handler2.CallsCount)

	event2 := TestEvent{Name: "AnotherEventName", Payload: "AnotherEventPayload"}
	err = dispatcher.Dispatch(event2)
	assert.NoError(t, err, "Error when dispatch event")

	assert.Equal(t, 1, handler.CallsCount)
	assert.Equal(t, 1, handler2.CallsCount)

	err = dispatcher.Dispatch(event)
	assert.NoError(t, err, "Error when dispatch event")

	assert.Equal(t, 2, handler.CallsCount)
	assert.Equal(t, 2, handler2.CallsCount)
}
