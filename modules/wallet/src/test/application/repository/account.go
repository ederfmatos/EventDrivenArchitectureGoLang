package repository

import "wallet/src/main/domain/entity"

type FakeAccountRepository struct {
	accounts map[string]*entity.Account
}

func NewFakeAccountRepository() *FakeAccountRepository {
	return &FakeAccountRepository{accounts: make(map[string]*entity.Account)}
}

func (repository *FakeAccountRepository) Save(account *entity.Account) error {
	repository.accounts[account.ID] = account
	return nil
}

func (repository *FakeAccountRepository) FindByID(id string) (*entity.Account, error) {
	return repository.accounts[id], nil
}

func (repository *FakeAccountRepository) UpdateBalance(account *entity.Account) error {
	repository.accounts[account.ID] = account
	return nil
}
