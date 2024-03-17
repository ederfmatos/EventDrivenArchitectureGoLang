package repository

import (
	"EventDrivenArchitectureGoLang/src/main/application/repository"
)

type FakeUnitOfWork struct {
	accountRepository     repository.AccountRepository
	transactionRepository repository.TransactionRepository
}

func NewFakeUnitOfWork(accountRepository repository.AccountRepository, transactionRepository repository.TransactionRepository) *FakeUnitOfWork {
	return &FakeUnitOfWork{accountRepository, transactionRepository}
}

func (unitOfWork *FakeUnitOfWork) Register(string, interface{}) {
}

func (unitOfWork *FakeUnitOfWork) GetRepository(name string) (interface{}, error) {
	switch name {
	case "ACCOUNT":
		return unitOfWork.accountRepository, nil
	case "TRANSACTION":
		return unitOfWork.transactionRepository, nil
	}
	return nil, nil
}

func (unitOfWork *FakeUnitOfWork) Do(fn func() error) error {
	return fn()
}
