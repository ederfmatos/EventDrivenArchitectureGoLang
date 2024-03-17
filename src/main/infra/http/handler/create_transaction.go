package handler

import (
	"EventDrivenArchitectureGoLang/src/main/application/usecase/transaction"
	"EventDrivenArchitectureGoLang/src/main/infra/http"
	"context"
	"io"
)

type CreateTransactionHandler struct {
	CreateTransactionUseCase *transaction.CreateTransactionUseCase
	ctx                      context.Context
}

func NewCreateTransactionHandler(createTransactionUseCase *transaction.CreateTransactionUseCase, ctx context.Context) *CreateTransactionHandler {
	return &CreateTransactionHandler{CreateTransactionUseCase: createTransactionUseCase, ctx: ctx}
}

func (handler *CreateTransactionHandler) Handle(body io.ReadCloser, server http.Server) (any, error) {
	var input transaction.CreateTransactionInput
	err := server.ParseBody(body, input)
	if err != nil {
		return nil, err
	}
	return handler.CreateTransactionUseCase.Execute(handler.ctx, input)
}
