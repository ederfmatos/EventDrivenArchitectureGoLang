package handler

import (
	"EventDrivenArchitectureGoLang/src/main/application/usecase/customer"
	"EventDrivenArchitectureGoLang/src/main/infra/http"
	"io"
)

type CreateCustomerHandler struct {
	CreateCustomerUseCase *customer.CreateCustomerUseCase
}

func NewCreateCustomerHandler(createCustomerUseCase *customer.CreateCustomerUseCase) *CreateCustomerHandler {
	return &CreateCustomerHandler{CreateCustomerUseCase: createCustomerUseCase}
}

func (handler *CreateCustomerHandler) Handle(body io.ReadCloser, server http.Server) (any, error) {
	var input customer.CreateCustomerInput
	err := server.ParseBody(body, input)
	if err != nil {
		return nil, err
	}
	return handler.CreateCustomerUseCase.Execute(input)
}
