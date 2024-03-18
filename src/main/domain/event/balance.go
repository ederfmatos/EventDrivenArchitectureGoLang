package event

import "time"

type BalanceUpdated struct {
	Name          string
	AccountFromId string
	AccountToId   string
}

func NewBalanceUpdatedEvent(accountFromId, accountToId string) *BalanceUpdated {
	return &BalanceUpdated{
		Name:          "BalanceUpdated",
		AccountFromId: accountFromId,
		AccountToId:   accountToId,
	}
}

func (event *BalanceUpdated) GetId() string {
	return "BALANCE"
}

func (event *BalanceUpdated) GetName() string {
	return event.Name
}

func (event *BalanceUpdated) GetDateTime() time.Time {
	return time.Now()
}
