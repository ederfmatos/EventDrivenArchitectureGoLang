package repository

import "wallet/src/main/domain/entity"

type FakeCustomerRepository struct {
	customers map[string]*entity.Customer
}

func NewFakeCustomerRepository() *FakeCustomerRepository {
	return &FakeCustomerRepository{customers: make(map[string]*entity.Customer)}
}

func (repository *FakeCustomerRepository) Get(id string) (*entity.Customer, error) {
	return repository.customers[id], nil
}

func (repository *FakeCustomerRepository) Save(customer *entity.Customer) error {
	repository.customers[customer.ID] = customer
	return nil
}
