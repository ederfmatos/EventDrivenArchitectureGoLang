package main

import (
	"EventDrivenArchitectureGoLang/src/main/application/usecase/account"
	"EventDrivenArchitectureGoLang/src/main/application/usecase/customer"
	"EventDrivenArchitectureGoLang/src/main/application/usecase/transaction"
	"EventDrivenArchitectureGoLang/src/main/infra/events"
	"EventDrivenArchitectureGoLang/src/main/infra/http"
	"EventDrivenArchitectureGoLang/src/main/infra/http/handler"
	"EventDrivenArchitectureGoLang/src/main/infra/repository"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetEnvValue(name, defaultValue string) string {
	env := os.Getenv(name)
	if env != "" {
		return env
	}
	if defaultValue != "" {
		return defaultValue
	}
	panic(fmt.Errorf("%s must be set", name))
}

func main() {
	databaseHost := GetEnvValue("DATABASE_HOST", "localhost")
	databasePort := GetEnvValue("DATABASE_PORT", "3306")
	databaseName := GetEnvValue("DATABASE_NAME", "wallet")
	databaseUserName := GetEnvValue("DATABASE_USERNAME", "root")
	databasePassword := GetEnvValue("DATABASE_PASSWORD", "root")
	kafkaServer := GetEnvValue("KAFKA_SERVER", "localhost:29092")
	kafkaGroupId := GetEnvValue("KAFKA_GROUP_ID", "wallet")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", databaseUserName, databasePassword, databaseHost, databasePort, databaseName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	kafkaEventEmitter := events.NewKafkaEventEmitter(kafkaServer, kafkaGroupId)

	customerRepository := repository.NewDefaultCustomerRepository(database)
	accountRepository := repository.NewDefaultAccountRepository(database)
	transactionRepository := repository.NewDefaultTransactionRepository(database)

	ctx := context.Background()
	unitOfWork := repository.NewDefaultUnitOfWork(database)

	unitOfWork.Register("ACCOUNT", accountRepository)
	unitOfWork.Register("TRANSACTION", transactionRepository)

	createTransactionUseCase := transaction.NewCreateTransactionUseCase(unitOfWork, kafkaEventEmitter)
	createCustomerUseCase := customer.NewCreateCustomerUseCase(customerRepository)
	createAccountUseCase := account.NewCreateAccountUseCase(accountRepository, customerRepository)

	httpServer := http.NewMuxHttpServer()
	createCustomerHandler := handler.NewCreateCustomerHandler(createCustomerUseCase)
	createAccountHandler := handler.NewCreateAccountHandler(createAccountUseCase)
	createTransactionHandler := handler.NewCreateTransactionHandler(createTransactionUseCase, ctx)
	httpServer.Post("/customers", createCustomerHandler.Handle)
	httpServer.Post("/accounts", createAccountHandler.Handle)
	httpServer.Post("/transactions", createTransactionHandler.Handle)
	httpServer.Listen("8080")
}
