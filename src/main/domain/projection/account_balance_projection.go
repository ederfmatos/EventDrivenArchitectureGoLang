package projection

type AccountBalanceProjection struct {
	AccountId    string  `gorm:"primaryKey" json:"accountId,omitempty"`
	CustomerId   string  `json:"customerId,omitempty"`
	CustomerName string  `json:"customerName,omitempty"`
	Balance      float64 `json:"balance,omitempty"`
}

func NewAccountBalanceProjection(customerId, customerName, accountId string, balance float64) *AccountBalanceProjection {
	return &AccountBalanceProjection{CustomerId: customerId, CustomerName: customerName, Balance: balance, AccountId: accountId}
}
