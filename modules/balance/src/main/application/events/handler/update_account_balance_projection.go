package handler

import (
	"balance/src/main/application/usecase/account"
	"balance/src/main/domain/event"
	"encoding/json"
)

type UpdateAccountBalanceProjectionEventHandler struct {
	*account.UpdateAccountBalanceProjectionUseCase
}

func NewUpdateAccountFromBalanceEventHandler(updateAccountBalanceProjectionUseCase *account.UpdateAccountBalanceProjectionUseCase) *UpdateAccountBalanceProjectionEventHandler {
	return &UpdateAccountBalanceProjectionEventHandler{UpdateAccountBalanceProjectionUseCase: updateAccountBalanceProjectionUseCase}
}

func (handler UpdateAccountBalanceProjectionEventHandler) EventName() string {
	return "BalanceUpdated"
}

func (handler UpdateAccountBalanceProjectionEventHandler) Handle(message []byte) error {
	var balanceUpdatedEvent event.BalanceUpdated
	err := json.Unmarshal(message, &balanceUpdatedEvent)
	if err != nil {
		return err
	}
	input := account.UpdateAccountBalanceProjectionInput{
		AccountId: balanceUpdatedEvent.Payload.AccountIdTo,
		Balance:   balanceUpdatedEvent.Payload.BalanceAccountIdTo,
	}
	err = handler.UpdateAccountBalanceProjectionUseCase.Execute(input)
	if err != nil {
		return err
	}
	input = account.UpdateAccountBalanceProjectionInput{
		AccountId: balanceUpdatedEvent.Payload.AccountIdFrom,
		Balance:   balanceUpdatedEvent.Payload.BalanceAccountIdFrom,
	}
	return handler.UpdateAccountBalanceProjectionUseCase.Execute(input)
}
