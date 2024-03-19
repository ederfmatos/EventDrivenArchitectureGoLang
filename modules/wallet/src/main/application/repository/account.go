package repository

import "wallet/src/main/domain/entity"

type AccountRepository interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
