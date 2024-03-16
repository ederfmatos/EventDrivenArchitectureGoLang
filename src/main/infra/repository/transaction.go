package repository

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"database/sql"
)

type DefaultTransactionRepository struct {
	DB *sql.DB
}

func NewDefaultTransactionRepository(db *sql.DB) *DefaultTransactionRepository {
	return &DefaultTransactionRepository{DB: db}
}

func (repository *DefaultTransactionRepository) Create(transaction *entity.Transaction) error {
	stmt, err := repository.DB.Prepare("INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID, transaction.Amount, transaction.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
