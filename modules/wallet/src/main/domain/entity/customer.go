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
	if customer.Name == "" {
		return nil, errors.New("name is required")
	}
	if customer.Email == "" {
		return nil, errors.New("email is required")
	}
	return customer, nil
}
