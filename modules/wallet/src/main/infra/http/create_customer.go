package http

import (
	"io"
	"net/http"
	"wallet/src/main/application/usecase/customer"
)

type CreateCustomerHandler struct {
	CreateCustomerUseCase *customer.CreateCustomerUseCase
}

func NewCreateCustomerHandler(createCustomerUseCase *customer.CreateCustomerUseCase) ServerHttpHandler {
	return &CreateCustomerHandler{CreateCustomerUseCase: createCustomerUseCase}
}

func (handler *CreateCustomerHandler) Handle(body io.ReadCloser, server Server, _ *http.Request) (any, error) {
	var input customer.CreateCustomerInput
	err := server.ParseBody(body, &input)
	if err != nil {
		return nil, err
	}
	return handler.CreateCustomerUseCase.Execute(input)
}
