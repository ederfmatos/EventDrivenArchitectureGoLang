package entity

import (
	errors2 "EventDrivenArchitectureGoLang/src/main/domain/errors"
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
		return nil, errors2.AmountMustBeGreaterThanZeroError
	}
	if accountFrom == nil {
		return nil, errors2.AccountFromNotFound
	}
	if accountTo == nil {
		return nil, errors2.AccountToNotFound
	}
	if accountFrom.Balance < amount {
		return nil, errors2.InsufficientFundError
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

func (t *Transaction) Commit() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}
