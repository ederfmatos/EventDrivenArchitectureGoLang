package customer

import (
	"EventDrivenArchitectureGoLang/src/main/application/repository"
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"time"
)

type CreateCustomerInput struct {
	Name  string
	Email string
}

type CreateCustomerOutput struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateCustomerUseCase struct {
	CustomerRepository repository.CustomerRepository
}

func NewCreateCustomerUseCase(customerRepository repository.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		CustomerRepository: customerRepository,
	}
}

func (useCase *CreateCustomerUseCase) Execute(input CreateCustomerInput) (*CreateCustomerOutput, error) {
	customer, err := entity.NewCustomer(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	err = useCase.CustomerRepository.Save(customer)
	if err != nil {
		return nil, err
	}
	output := &CreateCustomerOutput{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
	return output, nil
}
