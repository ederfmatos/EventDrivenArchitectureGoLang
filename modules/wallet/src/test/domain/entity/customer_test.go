package entity

import (
	"testing"
	"wallet/src/main/domain/entity"

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
