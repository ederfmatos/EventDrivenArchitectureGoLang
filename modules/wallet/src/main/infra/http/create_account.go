package http

import (
	"io"
	"net/http"
	"wallet/src/main/application/usecase/account"
)

type CreateAccountHandler struct {
	CreateAccountUseCase *account.CreateAccountUseCase
}

func NewCreateAccountHandler(createAccountUseCase *account.CreateAccountUseCase) *CreateAccountHandler {
	return &CreateAccountHandler{CreateAccountUseCase: createAccountUseCase}
}

func (handler *CreateAccountHandler) Handle(body io.ReadCloser, server Server, _ *http.Request) (any, error) {
	var input account.CreateAccountInput
	err := server.ParseBody(body, &input)
	if err != nil {
		return nil, err
	}
	return handler.CreateAccountUseCase.Execute(input)
}
