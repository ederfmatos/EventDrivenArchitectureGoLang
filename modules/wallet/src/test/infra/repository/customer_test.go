package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"wallet/src/main/domain/entity"
	"wallet/src/main/infra/repository"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type DefaultCustomerRepositoryTestSuite struct {
	suite.Suite
	db                        *gorm.DB
	defaultCustomerRepository *repository.DefaultCustomerRepository
}

func (suite *DefaultCustomerRepositoryTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	suite.Nil(err)
	suite.db = db
	//db.Exec("create table customers (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	suite.defaultCustomerRepository = repository.NewDefaultCustomerRepository(db)
}

func (suite *DefaultCustomerRepositoryTestSuite) TearDownSuite() {
	suite.db.Exec("DROP TABLE customers")
}

func TestDefaultCustomerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultCustomerRepositoryTestSuite))
}

func (suite *DefaultCustomerRepositoryTestSuite) TestSave() {
	customer := &entity.Customer{
		ID:    "1",
		Name:  "Test",
		Email: "j@j.com",
	}
	err := suite.defaultCustomerRepository.Save(customer)
	suite.Nil(err)
}

func (suite *DefaultCustomerRepositoryTestSuite) TestGet() {
	customer, _ := entity.NewCustomer("John", "j@j.com")
	_ = suite.defaultCustomerRepository.Save(customer)

	savedCustomer, err := suite.defaultCustomerRepository.Get(customer.ID)
	suite.Nil(err)
	suite.NotNil(savedCustomer)
	suite.Equal(customer.ID, savedCustomer.ID)
	suite.Equal(customer.Name, savedCustomer.Name)
	suite.Equal(customer.Email, savedCustomer.Email)
}
