package repository

import (
	"EventDrivenArchitectureGoLang/src/main/application/repository"
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/infra/repository/orm"
	"gorm.io/gorm"
)

type DefaultTransactionRepository struct {
	DB *gorm.DB
}

func NewDefaultTransactionRepository(DB *gorm.DB) repository.TransactionRepository {
	return &DefaultTransactionRepository{DB: DB}
}

func (repository *DefaultTransactionRepository) Create(transaction *entity.Transaction) error {
	return repository.DB.Create(orm.FromTransaction(transaction)).Error
}

func (repository *DefaultTransactionRepository) FindById(id string) (*entity.Transaction, error) {
	var transactionOrm orm.TransactionORM
	err := repository.DB.First(&transactionOrm, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	if transactionOrm.ID != "" {
		return transactionOrm.ToTransaction(), nil
	}
	return nil, nil
}
