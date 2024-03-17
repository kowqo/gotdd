package main

import (
	"errors"
	"fmt"
)

func main() {

	b := Bitcoin(300)
	fmt.Println(b)
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(num Bitcoin) {
	w.balance += num
}

func (w *Wallet) Balance() Bitcoin {

	return w.balance
}

var WithdrawError = errors.New("don't have bitcoin")

func (w *Wallet) Withdraw(b Bitcoin) error {

	if w.balance < b {
		return WithdrawError
	}

	w.balance -= b
	return nil

}
