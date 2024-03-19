package main

import (
	"balance/src/main/application/events/handler"
	"balance/src/main/application/usecase/account"
	"balance/src/main/infra"
	"balance/src/main/infra/events"
	"balance/src/main/infra/gateway"
	"balance/src/main/infra/http"
	"balance/src/main/infra/repository"
)

func main() {
	env := infra.NewEnv()
	database := repository.GetDatabase(env)

	accountBalanceProjectionGateway := gateway.NewDefaultAccountBalanceProjectionGateway(database)

	updateAccountBalanceProjectionUseCase := account.NewUpdateAccountBalanceProjectionUseCase(accountBalanceProjectionGateway)
	updateAccountBalanceProjectionEventHandler := handler.NewUpdateAccountFromBalanceEventHandler(updateAccountBalanceProjectionUseCase)
	transactionCreatedEventHandler := events.NewKafkaEventHandler(env.KafkaServer, env.KafkaGroupId)

	go transactionCreatedEventHandler.Consume(updateAccountBalanceProjectionEventHandler.EventName(), updateAccountBalanceProjectionEventHandler.Handle)

	httpServer := http.NewMuxHttpServer()
	httpServer.Get("/balances/{id}", http.NewGetAccountBalancesHandler(accountBalanceProjectionGateway))
	httpServer.Listen("3333")
}
