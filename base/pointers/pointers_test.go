package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(20))

		got := wallet.Balance()

		want := Bitcoin(20)

		if got != want {
			t.Errorf("%v got %s, but want %s", wallet, got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(120))
		wallet.Withdraw(Bitcoin(100))

		got := wallet.Balance()

		want := Bitcoin(20)

		if got != want {
			t.Errorf("%v got %s, but want %s", wallet, got, want)
		}
	})
}
