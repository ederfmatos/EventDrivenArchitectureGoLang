package http

import (
	"balance/src/main/application/gateway"
	"errors"
	"io"
	"net/http"
)

type GetAccountBalancesHandler struct {
	gateway.AccountBalanceProjectionGateway
}

func NewGetAccountBalancesHandler(accountBalanceProjectionGateway gateway.AccountBalanceProjectionGateway) *GetAccountBalancesHandler {
	return &GetAccountBalancesHandler{AccountBalanceProjectionGateway: accountBalanceProjectionGateway}
}

func (handler *GetAccountBalancesHandler) Handle(_ io.ReadCloser, _ Server, request *http.Request) (any, error) {
	id := request.PathValue("id")
	projection, err := handler.AccountBalanceProjectionGateway.FindByAccountId(id)
	if err != nil {
		return nil, err
	}
	if projection == nil {
		return nil, errors.New("account not found")
	}
	return projection, nil
}
