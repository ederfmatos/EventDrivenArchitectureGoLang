package transaction

import (
	"EventDrivenArchitectureGoLang/src/main/application/events"
	"EventDrivenArchitectureGoLang/src/main/application/repository"
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/domain/event"
	"context"
)

type CreateTransactionInput struct {
	AccountIdFrom string  `json:"account_id_from"`
	AccountIdTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutput struct {
	ID            string  `json:"id"`
	AccountIdFrom string  `json:"account_id_from"`
	AccountIdTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type BalanceUpdatedOutput struct {
	AccountIdFrom        string  `json:"account_id_from"`
	AccountIdTo          string  `json:"account_id_to"`
	BalanceAccountIdFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIdTo   float64 `json:"balance_account_id_to"`
}

type CreateTransactionUseCase struct {
	UnitOfWork      repository.UnitOfWork
	EventDispatcher events.EventDispatcher
}

func NewCreateTransactionUseCase(
	unitOfWork repository.UnitOfWork,
	eventDispatcher events.EventDispatcher,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		UnitOfWork:      unitOfWork,
		EventDispatcher: eventDispatcher,
	}
}

func (useCase *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInput) (*CreateTransactionOutput, error) {
	output := &CreateTransactionOutput{}
	balanceUpdatedOutput := &BalanceUpdatedOutput{}
	err := useCase.UnitOfWork.Do(ctx, func(_ *repository.UnitOfWork) error {
		accountRepository := useCase.getAccountRepository(ctx)
		transactionRepository := useCase.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindByID(input.AccountIdFrom)
		if err != nil {
			return err
		}
		accountTo, err := accountRepository.FindByID(input.AccountIdTo)
		if err != nil {
			return err
		}
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}
		output.ID = transaction.ID
		output.AccountIdFrom = input.AccountIdFrom
		output.AccountIdTo = input.AccountIdTo
		output.Amount = input.Amount

		balanceUpdatedOutput.AccountIdFrom = input.AccountIdFrom
		balanceUpdatedOutput.AccountIdTo = input.AccountIdTo
		balanceUpdatedOutput.BalanceAccountIdFrom = accountFrom.Balance
		balanceUpdatedOutput.BalanceAccountIdTo = accountTo.Balance
		return nil
	})
	if err != nil {
		return nil, err
	}

	transactionCreatedEvent := event.NewTransactionCreatedEvent(output)
	err = useCase.EventDispatcher.Dispatch(transactionCreatedEvent)
	if err != nil {
		return nil, err
	}

	balanceUpdatedEvent := event.NewBalanceUpdatedEvent(balanceUpdatedOutput)
	err = useCase.EventDispatcher.Dispatch(balanceUpdatedEvent)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (useCase *CreateTransactionUseCase) getAccountRepository(ctx context.Context) repository.AccountRepository {
	repo, err := useCase.UnitOfWork.GetRepository(ctx, "ACCOUNT")
	if err != nil {
		panic(err)
	}
	return repo.(repository.AccountRepository)
}

func (useCase *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) repository.TransactionRepository {
	repo, err := useCase.UnitOfWork.GetRepository(ctx, "TRANSACTION")
	if err != nil {
		panic(err)
	}
	return repo.(repository.TransactionRepository)
}
