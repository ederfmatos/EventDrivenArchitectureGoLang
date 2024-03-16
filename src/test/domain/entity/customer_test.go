package entity

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewCustomer(t *testing.T) {
	customer, err := entity.NewCustomer("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John Doe", customer.Name)
	assert.Equal(t, "j@j.com", customer.Email)
}

func TestCreateNewCustomerWhenArgsAreInvalid(t *testing.T) {
	customer, err := entity.NewCustomer("", "")
	assert.NotNil(t, err)
	assert.Nil(t, customer)
}

func TestUpdateCustomer(t *testing.T) {
	customer, _ := entity.NewCustomer("John Doe", "j@j.com")
	err := customer.Update("John Doe Update", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", customer.Name)
	assert.Equal(t, "j@j.com", customer.Email)
}

func TestUpdateCustomerWithInvalidArgs(t *testing.T) {
	customer, _ := entity.NewCustomer("John Doe", "j@j.com")
	err := customer.Update("", "j@j.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToCustomer(t *testing.T) {
	customer, _ := entity.NewCustomer("John Doe", "j@j")
	account := entity.NewAccount(customer)
	err := customer.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(customer.Accounts))
}
