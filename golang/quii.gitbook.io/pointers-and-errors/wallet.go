package pointersanderrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidAmount     = errors.New("invalid amount, must be positive or 0")
)

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount < 0 {
		return ErrInvalidAmount
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
