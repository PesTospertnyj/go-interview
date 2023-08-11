package bank

type Bank interface {
	Deposit(amount int) error

	Balance() (int, error)

	Close() (int, error)
}
