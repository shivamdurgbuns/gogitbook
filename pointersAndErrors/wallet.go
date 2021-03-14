package main

import (
	"errors"
	"fmt"
)

type Shivcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Shivcoin
}

func (s Shivcoin) String() string {
	return fmt.Sprintf("%d SHC", s)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

/*
Deposit method will decposti the money in the wallet.
*/
func (w *Wallet) Deposit(amount Shivcoin) {
	w.balance += amount
}

/*
Balance methhod will show the current balance of the wallet.
*/
func (w *Wallet) Balance() Shivcoin {
	return w.balance
}

/*
Withdraw method will withdraw coins from the current balance of the wallet.
*/
func (w *Wallet) Withdraw(amount Shivcoin) error {
	if amount < w.balance {
		w.balance -= amount
		return nil
	} else {
		return ErrInsufficientFunds
	}
}
