package main

type Wallet struct {
	balance int
}

/*
Deposit method will decposti the money in the wallet.
*/
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

/*
Balance methhod will show the current balance of the wallet.
*/
func (w *Wallet) Balance() int {
	return w.balance
}
