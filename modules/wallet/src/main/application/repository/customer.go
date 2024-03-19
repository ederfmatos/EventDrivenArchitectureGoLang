package repository

import "wallet/src/main/domain/entity"

type CustomerRepository interface {
	Get(id string) (*entity.Customer, error)
	Save(customer *entity.Customer) error
}
