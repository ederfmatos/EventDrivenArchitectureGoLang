package orm

import (
	"time"
	"wallet/src/main/domain/entity"
)

type TransactionORM struct {
	ID            string    `gorm:"id,omitempty"`
	AccountFromID string    `gorm:"account_from_id,omitempty"`
	AccountToID   string    `gorm:"account_to_id,omitempty"`
	Amount        float64   `gorm:"amount,omitempty"`
	CreatedAt     time.Time `gorm:"created_at,omitempty"`
}

func FromTransaction(transaction *entity.Transaction) *TransactionORM {
	return &TransactionORM{
		ID:            transaction.ID,
		AccountFromID: transaction.AccountFromID,
		AccountToID:   transaction.AccountToID,
		Amount:        transaction.Amount,
		CreatedAt:     transaction.CreatedAt,
	}
}

func (orm TransactionORM) TableName() string {
	return "transactions"
}

func (orm TransactionORM) ToTransaction() *entity.Transaction {
	return &entity.Transaction{
		ID:            orm.ID,
		AccountFromID: orm.AccountFromID,
		AccountToID:   orm.AccountToID,
		Amount:        orm.Amount,
		CreatedAt:     orm.CreatedAt,
	}
}
