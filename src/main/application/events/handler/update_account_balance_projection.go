package handler

import (
	"EventDrivenArchitectureGoLang/src/main/application/usecase/account"
	"EventDrivenArchitectureGoLang/src/main/domain/event"
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
		AccountId: balanceUpdatedEvent.AccountFromId,
	}
	err = handler.UpdateAccountBalanceProjectionUseCase.Execute(input)
	if err != nil {
		return err
	}
	input = account.UpdateAccountBalanceProjectionInput{
		AccountId: balanceUpdatedEvent.AccountToId,
	}
	return handler.UpdateAccountBalanceProjectionUseCase.Execute(input)
}
