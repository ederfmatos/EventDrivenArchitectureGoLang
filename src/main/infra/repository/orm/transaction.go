package orm

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"time"
)

type TransactionORM struct {
	ID            string
	AccountFrom   *AccountORM `gorm:"foreignKey:AccountFromID"`
	AccountFromID string
	AccountTo     *AccountORM `gorm:"foreignKey:AccountToID"`
	AccountToID   string
	Amount        float64
	CreatedAt     time.Time
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
