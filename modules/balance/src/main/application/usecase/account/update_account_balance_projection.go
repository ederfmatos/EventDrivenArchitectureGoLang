package account

import (
	"balance/src/main/application/gateway"
	"balance/src/main/domain/projection"
	"github.com/rs/zerolog/log"
)

type UpdateAccountBalanceProjectionUseCase struct {
	gateway.AccountBalanceProjectionGateway
}

func NewUpdateAccountBalanceProjectionUseCase(
	accountBalanceProjectionGateway gateway.AccountBalanceProjectionGateway,
) *UpdateAccountBalanceProjectionUseCase {
	return &UpdateAccountBalanceProjectionUseCase{
		AccountBalanceProjectionGateway: accountBalanceProjectionGateway,
	}
}

type UpdateAccountBalanceProjectionInput struct {
	AccountId string
	Balance   float64
}

func (useCase UpdateAccountBalanceProjectionUseCase) Execute(input UpdateAccountBalanceProjectionInput) error {
	accountBalanceProjection := projection.NewAccountBalanceProjection(input.AccountId, input.Balance)
	err := useCase.AccountBalanceProjectionGateway.Update(accountBalanceProjection)
	if err != nil {
		return err
	}
	log.Info().Any("balance", accountBalanceProjection).Msg("Account Balance Projection Updated")
	return nil
}
