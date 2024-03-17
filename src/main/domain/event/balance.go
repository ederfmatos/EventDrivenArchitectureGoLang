package event

import "time"

type BalanceUpdated struct {
	Name    string
	Payload interface{}
}

func NewBalanceUpdatedEvent(payload interface{}) *BalanceUpdated {
	return &BalanceUpdated{
		Name:    "BalanceUpdated",
		Payload: payload,
	}
}

func (event *BalanceUpdated) GetId() string {
	return "BALANCE"
}

func (event *BalanceUpdated) GetName() string {
	return event.Name
}

func (event *BalanceUpdated) GetPayload() interface{} {
	return event.Payload
}

func (event *BalanceUpdated) GetDateTime() time.Time {
	return time.Now()
}
