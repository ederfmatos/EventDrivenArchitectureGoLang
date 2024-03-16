package entity

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	customer1, _ := entity.NewCustomer("John Doe", "j@j")
	account1 := entity.NewAccount(customer1)
	customer2, _ := entity.NewCustomer("John Doe 2", "j@j2")
	account2 := entity.NewAccount(customer2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := entity.NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1100.0, account2.Balance)
	assert.Equal(t, 900.0, account1.Balance)
}

func TestCreateTransactionWithInsufficientBalance(t *testing.T) {
	customer1, _ := entity.NewCustomer("John Doe", "j@j")
	account1 := entity.NewAccount(customer1)
	customer2, _ := entity.NewCustomer("John Doe 2", "j@j2")
	account2 := entity.NewAccount(customer2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := entity.NewTransaction(account1, account2, 2000)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account1.Balance)
}
