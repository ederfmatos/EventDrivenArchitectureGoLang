package entity

import (
	"EventDrivenArchitectureGoLang/src/main/domain/errors"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID            string
	AccountFromID string
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
	transaction := &Transaction{
		ID:            uuid.New().String(),
		AccountFromID: accountFrom.ID,
		AccountToID:   accountTo.ID,
		Amount:        amount,
		CreatedAt:     time.Now(),
	}
	err := accountFrom.Debit(transaction.Amount)
	if err != nil {
		return nil, err
	}
	accountTo.Credit(transaction.Amount)
	return transaction, err
}
