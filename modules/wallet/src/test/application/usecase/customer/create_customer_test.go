package customer

import (
	"testing"
	"wallet/src/main/application/usecase/customer"
	repository2 "wallet/src/test/application/repository"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomerUseCase_Execute(t *testing.T) {
	customerRepository := repository2.NewFakeCustomerRepository()
	createCustomerUseCase := customer.NewCreateCustomerUseCase(customerRepository)

	input := customer.CreateCustomerInput{
		Name:  "John Doe",
		Email: "j@j",
	}
	output, err := createCustomerUseCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "j@j", output.Email)

	savedCustomer, err := customerRepository.Get(output.ID)
	assert.Nil(t, err)
	assert.NotNil(t, savedCustomer)
	assert.Equal(t, "John Doe", savedCustomer.Name)
	assert.Equal(t, "j@j", savedCustomer.Email)
}
