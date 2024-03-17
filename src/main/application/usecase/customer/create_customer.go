package customer

import (
	"EventDrivenArchitectureGoLang/src/main/application/repository"
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"time"
)

type CreateCustomerInput struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type CreateCustomerOutput struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCustomerUseCase struct {
	customerRepository repository.CustomerRepository
}

func NewCreateCustomerUseCase(customerRepository repository.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		customerRepository: customerRepository,
	}
}

func (useCase *CreateCustomerUseCase) Execute(input CreateCustomerInput) (*CreateCustomerOutput, error) {
	customer, err := entity.NewCustomer(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	err = useCase.customerRepository.Save(customer)
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
