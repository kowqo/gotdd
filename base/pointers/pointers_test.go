package main

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(b testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("%v got %s, but want %s", w, got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't got one")
		}

		if !errors.Is(want, got) {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		assertBalance(t, wallet, Bitcoin(20))

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)

		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw with no balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, WithdrawError)

		assertBalance(t, wallet, Bitcoin(20))

	})

}
