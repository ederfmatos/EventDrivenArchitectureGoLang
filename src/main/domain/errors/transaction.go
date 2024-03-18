package errors

import "errors"

var AmountMustBeGreaterThanZeroError = errors.New("amount must be greater than zero")
var InsufficientFundError = errors.New("insufficient funds")
var TransactionNotFoundError = errors.New("transaction not found")
