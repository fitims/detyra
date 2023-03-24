package wallet

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

type walletMock struct {
	onDeposit  func() error
	onWithdraw func() error
	onSend     func() error
	onBalance  func() float64
}

func (m *walletMock) DepositMoney(amount float64) error {
	return m.onDeposit()
}

func (m *walletMock) WithdrawMoney(amount float64) error {
	return m.onWithdraw()
}

func (m *walletMock) SendMoney(destination Wallet, amount float64) error {
	return m.onSend()
}

func (m *walletMock) CheckBalance() float64 {
	return m.onBalance()
}

func (m *walletMock) Handle(handler WalletHandler) {}

func TestNewWallet(t *testing.T) {
	wallet := NewWallet()
	balance := wallet.CheckBalance()

	if balance != 0 {
		t.Error(fmt.Sprintf("Balance is wrong. Want: 0, GOT: %f", balance))
	}
}

func TestNewWalletWithBalance(t *testing.T) {
	testData := []float64{
		0, 5, 10.5, 100, 120.567, 200.30, 563.45,
	}

	for _, b := range testData {
		wallet := NewWalletWithBalance(b)
		balance := wallet.CheckBalance()

		if balance != b {
			t.Error(fmt.Sprintf("Balance is wrong. Want: %f, GOT: %f", b, balance))
		}
	}
}

func TestDefaultWallet_DepositMoney(t *testing.T) {
	testData := []struct {
		Amount   float64
		Expected float64
		Err      error
	}{
		{Amount: 5, Expected: 5, Err: nil},
		{Amount: 5, Expected: 10, Err: nil},
		{Amount: 2.5, Expected: 12.5, Err: nil},
		{Amount: 17.13, Expected: 29.63, Err: nil},
		{Amount: 100.10, Expected: 129.73, Err: nil},
		// invalid amount
		{Amount: -10, Expected: 129.73, Err: InvalidAmountErr},
		{Amount: -5, Expected: 129.73, Err: InvalidAmountErr},
		{Amount: -1.2, Expected: 129.73, Err: InvalidAmountErr},
		{Amount: -0.5, Expected: 129.73, Err: InvalidAmountErr},
	}

	wallet := NewWallet()

	for _, v := range testData {
		err := wallet.DepositMoney(v.Amount)
		if err != v.Err {
			t.Error(fmt.Sprintf("Error is wrong. Want: %v, GOT: %v", v.Err, err))
		}

		balance := wallet.CheckBalance()
		if balance != v.Expected {
			t.Error(fmt.Sprintf("Balance is wrong. Want: %f, GOT: %f", v.Expected, balance))
		}
	}
}

func TestDefaultWallet_WithdrawMoney(t *testing.T) {
	testData := []struct {
		Amount   float64
		Expected float64
		Err      error
	}{
		{Amount: 5, Expected: 195, Err: nil},
		{Amount: 5, Expected: 190, Err: nil},
		{Amount: 2.5, Expected: 187.5, Err: nil},
		{Amount: 17.13, Expected: 170.37, Err: nil},
		{Amount: 100.10, Expected: 70.27, Err: nil},
		// invalid amount
		{Amount: -10, Expected: 70.27, Err: InvalidAmountErr},
		{Amount: -5, Expected: 70.27, Err: InvalidAmountErr},
		{Amount: -1.2, Expected: 70.27, Err: InvalidAmountErr},
		{Amount: -0.5, Expected: 70.27, Err: InvalidAmountErr},
		// insufficient funds
		{Amount: 70.30, Expected: 70.27, Err: InsufficientFundsErr},
		{Amount: 100, Expected: 70.27, Err: InsufficientFundsErr},
		{Amount: 90, Expected: 70.27, Err: InsufficientFundsErr},
	}

	wallet := NewWalletWithBalance(200)

	for _, v := range testData {
		err := wallet.WithdrawMoney(v.Amount)
		if err != v.Err {
			t.Error(fmt.Sprintf("Error is wrong. Want: %v, GOT: %v", v.Err, err))
		}

		balance := wallet.CheckBalance()

		if math.Abs(balance-v.Expected) > Epsilon {
			t.Error(fmt.Sprintf("Balance is wrong. Want: %f, GOT: %f", v.Expected, balance))
		}
	}
}

func TestDefaultWallet_SendMoney(t *testing.T) {
	testData := []struct {
		Amount   float64
		Expected float64
		Target   Wallet
		Err      error
	}{
		{Amount: 5, Expected: 195, Target: NewWallet(), Err: nil},
		{Amount: 5, Expected: 190, Target: NewWallet(), Err: nil},
		{Amount: 2.5, Expected: 187.5, Target: NewWallet(), Err: nil},
		{Amount: 17.13, Expected: 170.37, Target: NewWallet(), Err: nil},
		{Amount: 100.10, Expected: 70.27, Target: NewWallet(), Err: nil},
		// invalid amount
		{Amount: -10, Expected: 70.27, Target: NewWallet(), Err: InvalidAmountErr},
		{Amount: -5, Expected: 70.27, Target: NewWallet(), Err: InvalidAmountErr},
		{Amount: -1.2, Expected: 70.27, Target: NewWallet(), Err: InvalidAmountErr},
		{Amount: -0.5, Expected: 70.27, Target: NewWallet(), Err: InvalidAmountErr},
		// insufficient funds
		{Amount: 70.30, Expected: 70.27, Target: NewWallet(), Err: InsufficientFundsErr},
		{Amount: 100, Expected: 70.27, Target: NewWallet(), Err: InsufficientFundsErr},
		{Amount: 90, Expected: 70.27, Target: NewWallet(), Err: InsufficientFundsErr},
		// invalid target
		{Amount: 10, Expected: 70.27, Target: nil, Err: DestinationErr},
		// transaction error
		{
			Amount:   10,
			Expected: 70.27,
			Target: &walletMock{
				onDeposit: func() error {
					return errors.New("error depositing money")
				},
			},
			Err: TransactionErr,
		},
	}

	wallet := NewWalletWithBalance(200)

	for _, v := range testData {
		err := wallet.SendMoney(v.Target, v.Amount)
		if err != v.Err {
			t.Error(fmt.Sprintf("Error is wrong. Want: %v, GOT: %v", v.Err, err))
		}

		balance := wallet.CheckBalance()

		if math.Abs(balance-v.Expected) > Epsilon {
			t.Error(fmt.Sprintf("Balance is wrong. Want: %f, GOT: %f", v.Expected, balance))
		}
	}
}

func TestDefaultWallet_Handle(t *testing.T) {
	isCalled := false
	handler := func(id string, balance float64) {
		isCalled = true
	}

	wallet := NewWallet()
	wallet.Handle(handler)

	if !isCalled {
		t.Error("The handler is not called")
	}
}
