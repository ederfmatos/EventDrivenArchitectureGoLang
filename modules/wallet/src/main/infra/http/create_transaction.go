package http

import (
	"io"
	"net/http"
	"wallet/src/main/application/usecase/transaction"
)

type CreateTransactionHandler struct {
	CreateTransactionUseCase *transaction.CreateTransactionUseCase
}

func NewCreateTransactionHandler(createTransactionUseCase *transaction.CreateTransactionUseCase) ServerHttpHandler {
	return &CreateTransactionHandler{CreateTransactionUseCase: createTransactionUseCase}
}

func (handler *CreateTransactionHandler) Handle(body io.ReadCloser, server Server, _ *http.Request) (any, error) {
	var input transaction.CreateTransactionInput
	err := server.ParseBody(body, &input)
	if err != nil {
		return nil, err
	}
	return handler.CreateTransactionUseCase.Execute(input)
}
