package orm

import (
	"time"
	"wallet/src/main/domain/entity"
)

type AccountORM struct {
	ID         string    `gorm:"primaryKey"`
	CustomerId string    `gorm:"customer_id,omitempty"`
	Balance    float64   `gorm:"balance,omitempty"`
	CreatedAt  time.Time `gorm:"created_at,omitempty"`
	UpdatedAt  time.Time `gorm:"updated_at,omitempty"`
}

func FromAccount(account *entity.Account) *AccountORM {
	return &AccountORM{
		ID:         account.ID,
		CustomerId: account.CustomerId,
		Balance:    account.Balance,
		CreatedAt:  account.CreatedAt,
		UpdatedAt:  account.UpdatedAt,
	}
}

func (orm AccountORM) ToAccount() *entity.Account {
	return &entity.Account{
		ID:         orm.ID,
		CustomerId: orm.CustomerId,
		Balance:    orm.Balance,
		CreatedAt:  orm.CreatedAt,
		UpdatedAt:  orm.UpdatedAt,
	}
}

func (orm AccountORM) TableName() string {
	return "accounts"
}
