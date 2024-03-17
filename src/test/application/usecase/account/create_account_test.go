package account

import (
	"EventDrivenArchitectureGoLang/src/main/application/usecase/account"
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/domain/errors"
	"EventDrivenArchitectureGoLang/src/test/application/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	customer, _ := entity.NewCustomer("John Doe", "j@j")
	customerRepository := repository.NewFakeCustomerRepository()
	_ = customerRepository.Save(customer)

	accountRepository := repository.NewFakeAccountRepository()

	createAccountUseCase := account.NewCreateAccountUseCase(accountRepository, customerRepository)
	input := account.CreateAccountInput{CustomerID: customer.ID}
	output, err := createAccountUseCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)

	savedAccount, err := accountRepository.FindByID(output.ID)
	assert.Nil(t, err)
	assert.NotNil(t, savedAccount)
	assert.Equal(t, customer.ID, savedAccount.CustomerId)
	assert.Equal(t, 0.0, savedAccount.Balance)
}

func Test_Should_Return_Customer_Not_Found_Error(t *testing.T) {
	customerRepository := repository.NewFakeCustomerRepository()
	accountRepository := repository.NewFakeAccountRepository()

	createAccountUseCase := account.NewCreateAccountUseCase(accountRepository, customerRepository)
	input := account.CreateAccountInput{CustomerID: "any id"}
	output, err := createAccountUseCase.Execute(input)
	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.ErrorIs(t, errors.CustomerNotFoundError, err)
}
