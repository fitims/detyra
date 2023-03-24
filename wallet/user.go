package wallet

// User struct encapsulates the information about a user
type User struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Wallet Wallet `json:"wallet"`
}

// NewUser creates a new user with the provided email and name, and creates a wallet for user with 0 balance
func NewUser(email, name string) *User {
	return &User{
		Email:  email,
		Name:   name,
		Wallet: NewWallet(),
	}
}

// DepositMoneyIntoWallet will deposit the amount into the wallet. The call is delegated to the Wallet.DepositMoney
// method in the wallet
func (usr *User) DepositMoneyIntoWallet(amount float64) error {
	return usr.Wallet.DepositMoney(amount)
}

// WithdrawMoneyFromWallet will withdraw the amount from the wallet. The call is delegated to the Wallet.WithdrawMoney
// method in the wallet
func (usr *User) WithdrawMoneyFromWallet(amount float64) error {
	return usr.Wallet.WithdrawMoney(amount)
}

// SendMoneyTo will send the amount to the target user. The call is delegated to the Wallet.SendMoney method
// in the wallet by providing the target's wallet
func (usr *User) SendMoneyTo(target *User, amount float64) error {
	return usr.Wallet.SendMoney(target.Wallet, amount)
}

// CheckWalletBalance will check the wallet's balance. The call is delegated to the Wallet.CheckBalance
// method in the wallet
func (usr *User) CheckWalletBalance() float64 {
	return usr.Wallet.CheckBalance()
}
