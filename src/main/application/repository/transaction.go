package repository

import "EventDrivenArchitectureGoLang/src/main/domain/entity"

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
}
