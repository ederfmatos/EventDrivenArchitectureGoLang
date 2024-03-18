package gateway

import (
	"EventDrivenArchitectureGoLang/src/main/domain/projection"
	"gorm.io/gorm"
)

type DefaultAccountBalanceProjectionGateway struct {
	DB *gorm.DB
}

func NewDefaultAccountBalanceProjectionGateway(DB *gorm.DB) *DefaultAccountBalanceProjectionGateway {
	return &DefaultAccountBalanceProjectionGateway{DB: DB}
}

func (gateway *DefaultAccountBalanceProjectionGateway) Update(projection *projection.AccountBalanceProjection) error {
	return gateway.DB.Save(projection).Error
}

func (gateway *DefaultAccountBalanceProjectionGateway) FindByAccountId(id string) (*projection.AccountBalanceProjection, error) {
	var accountBalanceProjection *projection.AccountBalanceProjection
	err := gateway.DB.First(&accountBalanceProjection, "account_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return accountBalanceProjection, nil
}
