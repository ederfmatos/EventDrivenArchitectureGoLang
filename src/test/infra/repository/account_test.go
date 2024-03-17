package repository

import (
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/main/infra/repository"
	"gorm.io/gorm"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
)

type DefaultAccountRepositoryTestSuite struct {
	suite.Suite
	db                       *gorm.DB
	defaultAccountRepository *repository.DefaultAccountRepository
	customer                 *entity.Customer
}

func (suite *DefaultAccountRepositoryTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	suite.Nil(err)
	suite.db = db
	suite.defaultAccountRepository = repository.NewDefaultAccountRepository(db)
	suite.customer, _ = entity.NewCustomer("John", "j@j.com")
}

func (suite *DefaultAccountRepositoryTestSuite) TearDownSuite() {
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
	suite.db.Exec("INSERT INTO customers (id, name, email, created_at) values (?, ?, ?, ?)",
		suite.customer.ID, suite.customer.Name, suite.customer.Email, suite.customer.CreatedAt,
	)
	account := entity.NewAccount(suite.customer)
	err := suite.defaultAccountRepository.Save(account)
	suite.Nil(err)
	savedAccount, err := suite.defaultAccountRepository.FindByID(account.ID)
	suite.Nil(err)
	suite.Equal(account.ID, savedAccount.ID)
	suite.Equal(account.Balance, savedAccount.Balance)
	suite.Equal(account.CustomerId, savedAccount.CustomerId)
}
