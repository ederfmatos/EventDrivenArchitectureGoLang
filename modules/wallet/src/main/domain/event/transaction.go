package event

import "time"

type TransactionCreated struct {
	Name          string
	TransactionId string
}

func NewTransactionCreatedEvent(transactionId string) *TransactionCreated {
	return &TransactionCreated{
		Name:          "TransactionCreated",
		TransactionId: transactionId,
	}
}

func (event *TransactionCreated) GetId() string {
	return "TRANSACTIONS"
}

func (event *TransactionCreated) GetName() string {
	return event.Name
}

func (event *TransactionCreated) GetDateTime() time.Time {
	return time.Now()
}
