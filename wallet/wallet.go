package wallet

import (
	"github.com/google/uuid"
)

const (
	// epsilon defines the margin of error when comparing float64 numbers
	Epsilon = 0.00000001
)

// WalletHandler is a handler function that can be passed to the method Handle in the Wallet. The wallet will provide its id
// and current balance to the handler, which in turn can decide to process the info further. Once case scenario would be
// to provide a function that stores the current balance of the wallet into a data store and can be called after every transaction.
type WalletHandler func(id string, balance float64)

// Wallet encapsulates the business logic that a user/customer can do with a wallet.
type Wallet interface {
	DepositMoney(amount float64) error
	WithdrawMoney(amount float64) error
	SendMoney(destination Wallet, amount float64) error
	CheckBalance() float64
	Handle(handler WalletHandler)
}

// defaultWallet is the default implementation fo the Wallet. Each wallet has an Id and Balance
type defaultWallet struct {
	Id      string  `json:"id"`
	Balance float64 `json:"balance"`
}

// NewWallet creates a new wallet with a new Id and a balance of 0
func NewWallet() Wallet {
	return &defaultWallet{
		Id:      uuid.New().String(),
		Balance: 0,
	}
}

// NewWalletWithBalance creates a new wallet with a new Id and a balance as specified
func NewWalletWithBalance(balance float64) Wallet {
	return &defaultWallet{
		Id:      uuid.New().String(),
		Balance: balance,
	}
}

// DepositMoney will add the amount provided to the current balance. It first checks if the amount provided
// is a valid value (greater than 0). If the value is less than 0, then InvalidAmountErr error is returned
// otherwise the current balance is incremented by the provided amount
func (wallet *defaultWallet) DepositMoney(amount float64) error {

	// check if the amount is valid ie. greater than 0
	if amount <= 0 {
		return InvalidAmountErr
	}

	// increase the balance for the amount provided
	wallet.Balance += amount
	return nil
}

// WithdrawMoney will remove the amount provided from the current balance. It first checks if the amount provided
// is a valid value (greater than 0). If the value is less than 0, then InvalidAmountErr error is returned. Then
// the balance is checked to see if the eis enough funds to withdraw. If there is insufficient funds on the wallet
// InsufficientFundsErr error is returned. Otherwise, the amount provided is deducted from the current balance.
func (wallet *defaultWallet) WithdrawMoney(amount float64) error {

	// check if the amount is valid ie. greater than 0
	if amount <= 0 {
		return InvalidAmountErr
	}

	// check if there are sufficient funds to withdraw
	if wallet.Balance < amount {
		return InsufficientFundsErr
	}

	// Deduct the amount from the balance
	wallet.Balance -= amount
	return nil
}

// SendMoney will send the amount provided to the destination wallet. It checks if the amount provided is valid. If the
// amount is not valid then it returns  InvalidAmountErr. then checks if there are sufficient funds to transfer. If there
// are no sufficient funds, then InsufficientFundsErr is returned. Lastly, it checks if the destination is valid, and if it is valid
// it tries to deposit the amount into destination wallet. If the transfer fails, then TransactionErr is returned.
// If all the checks pass and the transfer succeeds, then the amount is deducted from the balance.
func (wallet *defaultWallet) SendMoney(destination Wallet, amount float64) error {

	// check if the amount is valid ie. greater than 0
	if amount <= 0 {
		return InvalidAmountErr
	}

	// check if there are sufficient funds to transfer
	if wallet.Balance < amount {
		return InsufficientFundsErr
	}

	// check if the destination wallet is valid
	if destination == nil {
		return DestinationErr
	}

	// try and deposit the money to the destination wallet
	err := destination.DepositMoney(amount)
	if err != nil {
		return TransactionErr
	}

	// Deduct the amount from the balance
	wallet.Balance -= amount
	return nil
}

// CheckBalance returns the current balance of the wallet
func (wallet *defaultWallet) CheckBalance() float64 {
	return wallet.Balance
}

// Handle calls the WalletHandler by providing the wallet id and current balance
func (wallet *defaultWallet) Handle(handler WalletHandler) {
	if handler != nil {
		handler(wallet.Id, wallet.Balance)
	}
}
