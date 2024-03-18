package repository

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/infra/repository/orm"
	"gorm.io/gorm"
)

type DefaultAccountRepository struct {
	DB *gorm.DB
}

func NewDefaultAccountRepository(DB *gorm.DB) *DefaultAccountRepository {
	return &DefaultAccountRepository{DB: DB}
}

func (repository *DefaultAccountRepository) FindByID(id string) (*entity.Account, error) {
	var accountOrm orm.AccountORM
	err := repository.DB.First(&accountOrm, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	if accountOrm.ID != "" {
		return accountOrm.ToAccount(), nil
	}
	return nil, nil
}

func (repository *DefaultAccountRepository) Save(account *entity.Account) error {
	accountORM := orm.FromAccount(account)
	return repository.DB.Create(accountORM).Error
}

func (repository *DefaultAccountRepository) UpdateBalance(account *entity.Account) error {
	accountORM := orm.FromAccount(account)
	return repository.DB.Model(accountORM).Update("balance", account.Balance).Error
}
