package account

import (
	"EventDrivenArchitectureGoLang/src/main/application/repository"
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/domain/errors"
)

type CreateAccountInput struct {
	CustomerID string `json:"customer_id"`
}

type CreateAccountOutput struct {
	ID string `json:"id,omitempty"`
}

type CreateAccountUseCase struct {
	AccountRepository  repository.AccountRepository
	CustomerRepository repository.CustomerRepository
}

func NewCreateAccountUseCase(AccountRepository repository.AccountRepository, CustomerRepository repository.CustomerRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountRepository:  AccountRepository,
		CustomerRepository: CustomerRepository,
	}
}

func (useCase *CreateAccountUseCase) Execute(input CreateAccountInput) (*CreateAccountOutput, error) {
	customer, err := useCase.CustomerRepository.Get(input.CustomerID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.CustomerNotFoundError
	}
	account := entity.NewAccount(customer)
	err = useCase.AccountRepository.Save(account)
	if err != nil {
		return nil, err
	}
	output := &CreateAccountOutput{ID: account.ID}
	return output, nil
}
