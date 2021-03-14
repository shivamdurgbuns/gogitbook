package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Shivcoin(10))
		got := wallet.Balance()
		want := Shivcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Shivcoin(20)}

		wallet.Withdraw(Shivcoin(10))
		got := wallet.Balance()
		want := Shivcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
