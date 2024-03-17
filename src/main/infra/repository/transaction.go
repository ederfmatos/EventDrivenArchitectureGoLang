package repository

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/infra/repository/orm"
	"gorm.io/gorm"
)

type DefaultTransactionRepository struct {
	DB *gorm.DB
}

func NewDefaultTransactionRepository(DB *gorm.DB) *DefaultTransactionRepository {
	_ = DB.AutoMigrate(orm.TransactionORM{})
	return &DefaultTransactionRepository{DB: DB}
}

func (repository *DefaultTransactionRepository) Create(transaction *entity.Transaction) error {
	return repository.DB.Create(orm.FromTransaction(transaction)).Error
}
