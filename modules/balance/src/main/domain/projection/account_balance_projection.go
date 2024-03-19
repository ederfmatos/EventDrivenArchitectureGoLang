package projection

type AccountBalanceProjection struct {
	AccountId string  `gorm:"primaryKey" json:"accountId,omitempty"`
	Balance   float64 `json:"balance,omitempty"`
}

func NewAccountBalanceProjection(accountId string, balance float64) *AccountBalanceProjection {
	return &AccountBalanceProjection{Balance: balance, AccountId: accountId}
}
