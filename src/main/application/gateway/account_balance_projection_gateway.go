package gateway

import "EventDrivenArchitectureGoLang/src/main/domain/projection"

type AccountBalanceProjectionGateway interface {
	Update(projection *projection.AccountBalanceProjection) error
	FindByAccountId(id string) (*projection.AccountBalanceProjection, error)
}
