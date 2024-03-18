package repository

import (
	"EventDrivenArchitectureGoLang/src/main/application/repository"
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/infra/repository/orm"
	"gorm.io/gorm"
)

type DefaultCustomerRepository struct {
	DB *gorm.DB
}

func NewDefaultCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &DefaultCustomerRepository{DB: db}
}

func (repository *DefaultCustomerRepository) Get(id string) (*entity.Customer, error) {
	var customerORM *orm.CustomerORM
	err := repository.DB.First(&customerORM, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	if customerORM.ID != "" {
		return customerORM.ToCustomer(), nil
	}
	return nil, nil
}

func (repository *DefaultCustomerRepository) Save(customer *entity.Customer) error {
	customerORM := orm.FromCustomer(customer)
	return repository.DB.Create(customerORM).Error
}
