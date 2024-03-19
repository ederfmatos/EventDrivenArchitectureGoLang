package main

import (
	"wallet/src/main/application/usecase/account"
	"wallet/src/main/application/usecase/customer"
	"wallet/src/main/application/usecase/transaction"
	"wallet/src/main/infra"
	"wallet/src/main/infra/events"
	"wallet/src/main/infra/http"
	"wallet/src/main/infra/repository"
)

func main() {
	env := infra.NewEnv()
	database := repository.GetDatabase(env)

	kafkaEventEmitter := events.NewKafkaEventEmitter(env.KafkaServer)
	customerRepository := repository.NewDefaultCustomerRepository(database)
	accountRepository := repository.NewDefaultAccountRepository(database)
	transactionRepository := repository.NewDefaultTransactionRepository(database)

	unitOfWork := repository.NewDefaultUnitOfWork(database)
	unitOfWork.Register("ACCOUNT", accountRepository)
	unitOfWork.Register("TRANSACTION", transactionRepository)

	createTransactionUseCase := transaction.NewCreateTransactionUseCase(unitOfWork, kafkaEventEmitter)
	createCustomerUseCase := customer.NewCreateCustomerUseCase(customerRepository)
	createAccountUseCase := account.NewCreateAccountUseCase(accountRepository, customerRepository)

	httpServer := http.NewMuxHttpServer()
	httpServer.Post("/customers", http.NewCreateCustomerHandler(createCustomerUseCase))
	httpServer.Post("/accounts", http.NewCreateAccountHandler(createAccountUseCase))
	httpServer.Post("/transactions", http.NewCreateTransactionHandler(createTransactionUseCase))
	httpServer.Listen("8080")
}
