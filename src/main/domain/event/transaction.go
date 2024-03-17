package event

import "time"

type TransactionCreated struct {
	Name    string
	Payload interface{}
}

func NewTransactionCreatedEvent(payload interface{}) *TransactionCreated {
	return &TransactionCreated{
		Name:    "TransactionCreated",
		Payload: payload,
	}
}

func (event *TransactionCreated) GetId() string {
	return "TRANSACTIONS"
}

func (event *TransactionCreated) GetName() string {
	return event.Name
}

func (event *TransactionCreated) GetPayload() interface{} {
	return event.Payload
}

func (event *TransactionCreated) GetDateTime() time.Time {
	return time.Now()
}
