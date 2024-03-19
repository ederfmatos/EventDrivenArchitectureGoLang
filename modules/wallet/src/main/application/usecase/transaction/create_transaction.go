package transaction

import (
	"wallet/src/main/application/events"
	"wallet/src/main/application/repository"
	"wallet/src/main/domain/entity"
	"wallet/src/main/domain/event"
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
	UnitOfWork   repository.UnitOfWork
	EventEmitter events.EventEmitter
}

func NewCreateTransactionUseCase(
	unitOfWork repository.UnitOfWork,
	eventEmitter events.EventEmitter,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		UnitOfWork:   unitOfWork,
		EventEmitter: eventEmitter,
	}
}

func (useCase *CreateTransactionUseCase) Execute(input CreateTransactionInput) (*CreateTransactionOutput, error) {
	output := &CreateTransactionOutput{}
	balanceUpdatedOutput := &BalanceUpdatedOutput{}
	err := useCase.UnitOfWork.Do(func() error {
		accountRepository := useCase.getAccountRepository()
		transactionRepository := useCase.getTransactionRepository()

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

	transactionCreatedEvent := event.NewTransactionCreatedEvent(output.ID)
	err = useCase.EventEmitter.Emit(transactionCreatedEvent)
	if err != nil {
		return nil, err
	}

	balanceUpdatedEvent := event.NewBalanceUpdatedEvent(balanceUpdatedOutput)
	err = useCase.EventEmitter.Emit(balanceUpdatedEvent)

	if err != nil {
		return nil, err
	}
	return output, nil
}

func (useCase *CreateTransactionUseCase) getAccountRepository() repository.AccountRepository {
	repo, err := useCase.UnitOfWork.GetRepository("ACCOUNT")
	if err != nil {
		panic(err)
	}
	return repo.(repository.AccountRepository)
}

func (useCase *CreateTransactionUseCase) getTransactionRepository() repository.TransactionRepository {
	repo, err := useCase.UnitOfWork.GetRepository("TRANSACTION")
	if err != nil {
		panic(err)
	}
	return repo.(repository.TransactionRepository)
}
