package repository

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"database/sql"
)

type DefaultCustomerRepository struct {
	DB *sql.DB
}

func NewDefaultCustomerRepository(db *sql.DB) *DefaultCustomerRepository {
	return &DefaultCustomerRepository{DB: db}
}

func (repository *DefaultCustomerRepository) Get(id string) (*entity.Customer, error) {
	customer := &entity.Customer{}
	stmt, err := repository.DB.Prepare("SELECT id, name, email, created_at FROM customers WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	if err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt); err != nil {
		return nil, err
	}
	return customer, nil
}

func (repository *DefaultCustomerRepository) Save(customer *entity.Customer) error {
	stmt, err := repository.DB.Prepare("INSERT INTO customers (id, name, email, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email, customer.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
