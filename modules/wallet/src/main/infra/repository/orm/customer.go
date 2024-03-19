package orm

import (
	"time"
	"wallet/src/main/domain/entity"
)

type CustomerORM struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"name,omitempty"`
	Email     string    `gorm:"email,omitempty"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func FromCustomer(customer *entity.Customer) *CustomerORM {
	return &CustomerORM{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}

func (orm CustomerORM) ToCustomer() *entity.Customer {
	return &entity.Customer{
		ID:        orm.ID,
		Name:      orm.Name,
		Email:     orm.Email,
		CreatedAt: orm.CreatedAt,
		UpdatedAt: orm.UpdatedAt,
	}
}

func (orm CustomerORM) TableName() string {
	return "customers"
}
