package repository

import (
	"EventDrivenArchitectureGoLang/src/main/application/repository"
	"context"
)

type FakeUnitOfWork struct {
	accountRepository     repository.AccountRepository
	transactionRepository repository.TransactionRepository
}

func NewFakeUnitOfWork(accountRepository repository.AccountRepository, transactionRepository repository.TransactionRepository) *FakeUnitOfWork {
	return &FakeUnitOfWork{accountRepository, transactionRepository}
}

func (unitOfWork *FakeUnitOfWork) Register(string, repository.Factory) {
}

func (unitOfWork *FakeUnitOfWork) GetRepository(_ context.Context, name string) (interface{}, error) {
	switch name {
	case "ACCOUNT":
		return unitOfWork.accountRepository, nil
	case "TRANSACTION":
		return unitOfWork.transactionRepository, nil
	}
	return nil, nil
}

func (unitOfWork *FakeUnitOfWork) Do(ctx context.Context, fn func(unitOfWork *repository.UnitOfWork) error) error {
	return fn(nil)
}

func (unitOfWork *FakeUnitOfWork) CommitOrRollback() error {
	return nil
}

func (unitOfWork *FakeUnitOfWork) Rollback() error {
	return nil
}

func (unitOfWork *FakeUnitOfWork) UnRegister(name string) {
}
