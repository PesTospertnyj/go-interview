package bank

import "errors"

type Account struct {
	Amount   int
	isClosed bool
}

var (
	ErrAccountClosed  = errors.New("account is closed")
	ErrNotEnoughFunds = errors.New("not enough funds")
)

func NewBankAccount() Bank {
	return &Account{}
}

func (a *Account) Deposit(amount int) error {
	if a.isClosed {
		return ErrAccountClosed
	}

	if a.Amount+amount < 0 {
		return ErrNotEnoughFunds
	}

	a.Amount += amount

	return nil
}

func (a *Account) Balance() (int, error) {
	if a.isClosed {
		return 0, ErrAccountClosed
	}

	return a.Amount, nil
}

func (a *Account) Close() (int, error) {
	if a.isClosed {
		return 0, ErrAccountClosed
	}

	a.isClosed = !a.isClosed
	a.Amount = 0

	return 0, nil
}
