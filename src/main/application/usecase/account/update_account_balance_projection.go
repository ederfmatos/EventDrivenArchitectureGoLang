package account

import (
	"EventDrivenArchitectureGoLang/src/main/application/gateway"
	"EventDrivenArchitectureGoLang/src/main/application/repository"
	"EventDrivenArchitectureGoLang/src/main/domain/projection"
	"github.com/rs/zerolog/log"
)

type UpdateAccountBalanceProjectionUseCase struct {
	repository.TransactionRepository
	repository.CustomerRepository
	repository.AccountRepository
	gateway.AccountBalanceProjectionGateway
}

func NewUpdateAccountBalanceProjectionUseCase(
	transactionRepository repository.TransactionRepository,
	customerRepository repository.CustomerRepository,
	accountRepository repository.AccountRepository,
	accountBalanceProjectionGateway gateway.AccountBalanceProjectionGateway,
) *UpdateAccountBalanceProjectionUseCase {
	return &UpdateAccountBalanceProjectionUseCase{
		TransactionRepository:           transactionRepository,
		CustomerRepository:              customerRepository,
		AccountRepository:               accountRepository,
		AccountBalanceProjectionGateway: accountBalanceProjectionGateway,
	}
}

type UpdateAccountBalanceProjectionInput struct {
	AccountId string
}

func (useCase UpdateAccountBalanceProjectionUseCase) Execute(input UpdateAccountBalanceProjectionInput) error {
	account, err := useCase.AccountRepository.FindByID(input.AccountId)
	if err != nil {
		return err
	}
	customer, err := useCase.CustomerRepository.Get(account.CustomerId)
	if err != nil {
		return err
	}
	accountBalanceProjection := projection.NewAccountBalanceProjection(customer.ID, customer.Name, account.ID, account.Balance)
	err = useCase.AccountBalanceProjectionGateway.Update(accountBalanceProjection)
	if err != nil {
		return err
	}
	log.Info().Any("accountBalanceProjection", accountBalanceProjection).Msg("Account Balance Projection Updated")
	return nil
}
