package entity

import (
	"EventDrivenArchitectureGoLang/src/main/domain/errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID         string
	CustomerId string
	Balance    float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewAccount(customer *Customer) *Account {
	if customer == nil {
		return nil
	}
	account := &Account{
		ID:         uuid.New().String(),
		CustomerId: customer.ID,
		Balance:    0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return account
}

func (account *Account) Credit(amount float64) {
	account.Balance += amount
	account.UpdatedAt = time.Now()
}

func (account *Account) Debit(amount float64) error {
	if account.Balance < amount {
		return errors.InsufficientFundError
	}
	account.Balance -= amount
	account.UpdatedAt = time.Now()
	return nil
}
