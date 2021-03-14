package main

import "fmt"

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
Withdraw methhod will will withdraw coins from the current balance of the wallet.
*/
func (w *Wallet) Withdraw(amount Shivcoin) {
	if amount < w.balance {
		w.balance -= amount
	} else {
		return
	}
}
