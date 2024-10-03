package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw without insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(10))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if err != want {
		t.Errorf("got: %q, want: %q", err, want)
	}
}
