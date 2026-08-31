package pointerserrors

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance = w.balance + amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("cannot withdraw insufficient funds")
	}
	w.balance -= amount
	return nil
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
