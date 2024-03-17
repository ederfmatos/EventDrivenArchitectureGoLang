package entity

import (
	"EventDrivenArchitectureGoLang/src/main/domain/errors"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID            string
	AccountFrom   *Account `gorm:"foreignKey:AccountFromID"`
	AccountFromID string
	AccountTo     *Account `gorm:"foreignKey:AccountToID"`
	AccountToID   string
	Amount        float64
	CreatedAt     time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	if amount <= 0 {
		return nil, errors.AmountMustBeGreaterThanZeroError
	}
	if accountFrom == nil {
		return nil, errors.AccountFromNotFound
	}
	if accountTo == nil {
		return nil, errors.AccountToNotFound
	}
	if accountFrom.Balance < amount {
		return nil, errors.InsufficientFundError
	}
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	transaction.Commit()
	return transaction, nil
}

func (transaction *Transaction) Commit() {
	transaction.AccountFrom.Debit(transaction.Amount)
	transaction.AccountTo.Credit(transaction.Amount)
}
