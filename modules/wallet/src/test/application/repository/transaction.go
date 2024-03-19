package repository

import "wallet/src/main/domain/entity"

type FakeTransactionRepository struct {
	transactions map[string]*entity.Transaction
}

func NewFakeTransactionRepository() *FakeTransactionRepository {
	return &FakeTransactionRepository{transactions: make(map[string]*entity.Transaction)}
}

func (repository *FakeTransactionRepository) Create(transaction *entity.Transaction) error {
	repository.transactions[transaction.ID] = transaction
	return nil
}
