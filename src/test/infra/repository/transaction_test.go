package repository

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/infra/repository"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type DefaultTransactionRepositoryTestSuite struct {
	suite.Suite
	db                           *sql.DB
	customer                     *entity.Customer
	customer2                    *entity.Customer
	accountFrom                  *entity.Account
	accountTo                    *entity.Account
	defaultTransactionRepository *repository.DefaultTransactionRepository
}

func (s *DefaultTransactionRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("create table customers (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("create table accounts (id varchar(255), customer_id varchar(255), balance int, created_at date)")
	db.Exec("create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	customer, err := entity.NewCustomer("John", "j@j.com")
	s.Nil(err)
	s.customer = customer
	customer2, err := entity.NewCustomer("John2", "jj@j.com")
	s.Nil(err)
	s.customer2 = customer2
	//creating accounts
	accountFrom := entity.NewAccount(s.customer)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom
	accountTo := entity.NewAccount(s.customer2)
	accountTo.Balance = 1000
	s.accountTo = accountTo
	s.defaultTransactionRepository = repository.NewDefaultTransactionRepository(db)
}

func (s *DefaultTransactionRepositoryTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE customers")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestDefaultTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultTransactionRepositoryTestSuite))
}

func (s *DefaultTransactionRepositoryTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.defaultTransactionRepository.Create(transaction)
	s.Nil(err)
}
