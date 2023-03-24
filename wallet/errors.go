package wallet

import "errors"

// wallet errors
var (
	InvalidAmountErr     = errors.New("invalid amount")
	InsufficientFundsErr = errors.New("wallet has insufficient funds")
	DestinationErr       = errors.New("destination wallet is not valid")
	TransactionErr       = errors.New("transaction is not valid")
)

// user errors
var (
	UserAlreadyRegisteredErr = errors.New("user is already registered")
	UserDoesNotExistErr      = errors.New("user does not exist")
)
