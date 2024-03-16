package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name string, email string) (*Customer, error) {
	customer := &Customer{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := customer.Validate()
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (customer *Customer) Validate() error {
	if customer.Name == "" {
		return errors.New("name is required")
	}
	if customer.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (customer *Customer) Update(name string, email string) error {
	customer.Name = name
	customer.Email = email
	customer.UpdatedAt = time.Now()
	err := customer.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (customer *Customer) AddAccount(account *Account) error {
	if account.Customer.ID != customer.ID {
		return errors.New("account does not belong to customer")
	}
	customer.Accounts = append(customer.Accounts, account)
	return nil
}
