package main

import (
	"EventDrivenArchitectureGoLang/src/main/application/events/handler"
	"EventDrivenArchitectureGoLang/src/main/application/usecase/account"
	"EventDrivenArchitectureGoLang/src/main/application/usecase/customer"
	"EventDrivenArchitectureGoLang/src/main/application/usecase/transaction"
	"EventDrivenArchitectureGoLang/src/main/infra"
	"EventDrivenArchitectureGoLang/src/main/infra/events"
	"EventDrivenArchitectureGoLang/src/main/infra/gateway"
	"EventDrivenArchitectureGoLang/src/main/infra/http"
	"EventDrivenArchitectureGoLang/src/main/infra/repository"
)

func main() {
	env := infra.NewEnv()
	database := repository.GetDatabase(env)

	kafkaEventEmitter := events.NewKafkaEventEmitter(env.KafkaServer)
	customerRepository := repository.NewDefaultCustomerRepository(database)
	accountRepository := repository.NewDefaultAccountRepository(database)
	transactionRepository := repository.NewDefaultTransactionRepository(database)
	accountBalanceProjectionGateway := gateway.NewDefaultAccountBalanceProjectionGateway(database)

	unitOfWork := repository.NewDefaultUnitOfWork(database)
	unitOfWork.Register("ACCOUNT", accountRepository)
	unitOfWork.Register("TRANSACTION", transactionRepository)

	createTransactionUseCase := transaction.NewCreateTransactionUseCase(unitOfWork, kafkaEventEmitter)
	createCustomerUseCase := customer.NewCreateCustomerUseCase(customerRepository)
	createAccountUseCase := account.NewCreateAccountUseCase(accountRepository, customerRepository)
	updateAccountBalanceProjectionUseCase := account.NewUpdateAccountBalanceProjectionUseCase(transactionRepository, customerRepository, accountRepository, accountBalanceProjectionGateway)
	updateAccountBalanceProjectionEventHandler := handler.NewUpdateAccountFromBalanceEventHandler(updateAccountBalanceProjectionUseCase)
	transactionCreatedEventHandler := events.NewKafkaEventHandler(env.KafkaServer, env.KafkaGroupId)

	go transactionCreatedEventHandler.Consume(updateAccountBalanceProjectionEventHandler.EventName(), updateAccountBalanceProjectionEventHandler.Handle)

	httpServer := http.NewMuxHttpServer()
	httpServer.Get("/balances/{id}", http.NewGetAccountBalancesHandler(accountBalanceProjectionGateway))
	httpServer.Post("/customers", http.NewCreateCustomerHandler(createCustomerUseCase))
	httpServer.Post("/accounts", http.NewCreateAccountHandler(createAccountUseCase))
	httpServer.Post("/transactions", http.NewCreateTransactionHandler(createTransactionUseCase))
	httpServer.Listen("8080")
}
