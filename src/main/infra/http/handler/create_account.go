package handler

import (
	"EventDrivenArchitectureGoLang/src/main/application/usecase/account"
	"EventDrivenArchitectureGoLang/src/main/infra/http"
	"io"
)

type CreateAccountHandler struct {
	CreateAccountUseCase *account.CreateAccountUseCase
}

func NewCreateAccountHandler(createAccountUseCase *account.CreateAccountUseCase) *CreateAccountHandler {
	return &CreateAccountHandler{CreateAccountUseCase: createAccountUseCase}
}

func (handler *CreateAccountHandler) Handle(body io.ReadCloser, server http.Server) (any, error) {
	var input account.CreateAccountInput
	err := server.ParseBody(body, input)
	if err != nil {
		return nil, err
	}
	return handler.CreateAccountUseCase.Execute(input)
}
