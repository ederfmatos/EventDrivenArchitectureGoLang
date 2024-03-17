package transaction

import (
	"EventDrivenArchitectureGoLang/src/main/application/usecase/transaction"
	"EventDrivenArchitectureGoLang/src/main/domain/entity"
	"EventDrivenArchitectureGoLang/src/test/application/repository"
	events2 "EventDrivenArchitectureGoLang/src/test/infra/events"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	accountRepository := repository.NewFakeAccountRepository()
	transactionRepository := repository.NewFakeTransactionRepository()

	customer1, _ := entity.NewCustomer("customer1", "j@j.com")
	account1 := entity.NewAccount(customer1)
	account1.Credit(1000)
	_ = accountRepository.Save(account1)

	customer2, _ := entity.NewCustomer("customer2", "j@j2.com")
	account2 := entity.NewAccount(customer2)
	account2.Credit(1000)
	_ = accountRepository.Save(account2)

	unitOfWork := repository.NewFakeUnitOfWork(accountRepository, transactionRepository)

	input := transaction.CreateTransactionInput{
		AccountIdFrom: account1.ID,
		AccountIdTo:   account2.ID,
		Amount:        100,
	}

	emitter := events2.NewSpyEventEmitter()

	createTransactionUseCase := transaction.NewCreateTransactionUseCase(unitOfWork, emitter)
	output, err := createTransactionUseCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
}
