package entity

import (
	"testing"
	"wallet/src/main/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	customer, _ := entity.NewCustomer("John Doe", "j@j")
	account := entity.NewAccount(customer)
	assert.NotNil(t, account)
	assert.Equal(t, customer.ID, account.CustomerId)
}

func TestCreateAccountWithNilCustomer(t *testing.T) {
	account := entity.NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	customer, _ := entity.NewCustomer("John Doe", "j@j")
	account := entity.NewAccount(customer)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	customer, _ := entity.NewCustomer("John Doe", "j@j")
	account := entity.NewAccount(customer)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, float64(50), account.Balance)
}
