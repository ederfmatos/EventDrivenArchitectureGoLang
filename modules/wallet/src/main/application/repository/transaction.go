package repository

import "wallet/src/main/domain/entity"

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
	FindById(id string) (*entity.Transaction, error)
}
