package repository

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/infra/repository"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type DefaultAccountRepositoryTestSuite struct {
	suite.Suite
	db                       *sql.DB
	defaultAccountRepository *repository.DefaultAccountRepository
	customer                 *entity.Customer
}

func (suite *DefaultAccountRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("create table customers (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("create table accounts (id varchar(255), customer_id varchar(255), balance int, created_at date)")
	suite.defaultAccountRepository = repository.NewDefaultAccountRepository(db)
	suite.customer, _ = entity.NewCustomer("John", "j@j.com")
}

func (suite *DefaultAccountRepositoryTestSuite) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE customers")
	suite.db.Exec("DROP TABLE accounts")
}

func TestDefaultAccountRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultAccountRepositoryTestSuite))
}

func (suite *DefaultAccountRepositoryTestSuite) TestSave() {
	account := entity.NewAccount(suite.customer)
	err := suite.defaultAccountRepository.Save(account)
	suite.Nil(err)
}

func (suite *DefaultAccountRepositoryTestSuite) TestFindByID() {
	suite.db.Exec("Insert into customers (id, name, email, created_at) values (?, ?, ?, ?)",
		suite.customer.ID, suite.customer.Name, suite.customer.Email, suite.customer.CreatedAt,
	)
	account := entity.NewAccount(suite.customer)
	err := suite.defaultAccountRepository.Save(account)
	suite.Nil(err)
	savedAccount, err := suite.defaultAccountRepository.FindByID(account.ID)
	suite.Nil(err)
	suite.Equal(account.ID, savedAccount.ID)
	suite.Equal(account.Balance, savedAccount.Balance)
	suite.Equal(account.Customer.ID, savedAccount.Customer.ID)
	suite.Equal(account.Customer.Name, savedAccount.Customer.Name)
	suite.Equal(account.Customer.Email, savedAccount.Customer.Email)
}
