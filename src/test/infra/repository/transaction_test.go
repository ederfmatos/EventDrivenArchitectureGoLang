package repository

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/infra/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type DefaultTransactionRepositoryTestSuite struct {
	suite.Suite
	db                           *gorm.DB
	customer                     *entity.Customer
	customer2                    *entity.Customer
	accountFrom                  *entity.Account
	accountTo                    *entity.Account
	defaultTransactionRepository *repository.DefaultTransactionRepository
}

func (suite *DefaultTransactionRepositoryTestSuite) SetupSuite() {
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: gormLogger})
	suite.Nil(err)
	suite.db = db
	db.Exec("create table customers (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("create table accounts (id varchar(255), customer_id varchar(255), balance int, created_at date)")
	db.Exec("create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	customer, err := entity.NewCustomer("John", "j@j.com")
	suite.Nil(err)
	suite.customer = customer
	customer2, err := entity.NewCustomer("John2", "jj@j.com")
	suite.Nil(err)
	suite.customer2 = customer2
	//creating accounts
	accountFrom := entity.NewAccount(suite.customer)
	accountFrom.Balance = 1000
	suite.accountFrom = accountFrom
	accountTo := entity.NewAccount(suite.customer2)
	accountTo.Balance = 1000
	suite.accountTo = accountTo
	suite.defaultTransactionRepository = repository.NewDefaultTransactionRepository(db)
}

func (suite *DefaultTransactionRepositoryTestSuite) TearDownSuite() {
	suite.db.Exec("DROP TABLE customers")
	suite.db.Exec("DROP TABLE accounts")
	suite.db.Exec("DROP TABLE transactions")
}

func TestDefaultTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultTransactionRepositoryTestSuite))
}

func (suite *DefaultTransactionRepositoryTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(suite.accountFrom, suite.accountTo, 100)
	suite.Nil(err)
	err = suite.defaultTransactionRepository.Create(transaction)
	suite.Nil(err)
}
