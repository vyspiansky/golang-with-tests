package main

import (
	"errors"
	"fmt"
)

// Create new types from existing ones
type Bitcoin int

// Define how your type is printed when used with
// the %s format string in prints
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// if a symbol (variables, types, functions) starts with a lowercase symbol
	// then it is private
	balance Bitcoin
}

// The *Wallet you can read as "a pointer to a wallet".
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %p \n", &w.balance)

	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// func (w *Wallet) Withdraw(amount Bitcoin) {
// 	w.balance -= amount
// }

// func (w *Wallet) Withdraw(amount Bitcoin) error {
// 	w.balance -= amount
// 	return nil
// }

// func (w *Wallet) Withdraw(amount Bitcoin) error {
// 	if amount > w.balance {
// 		// errors.New creates a new error with a message of your choosing.
// 		return errors.New("oh no")
// 	}

// 	w.balance -= amount
// 	return nil
// }

// func (w *Wallet) Withdraw(amount Bitcoin) error {
// 	if amount > w.balance {
// 		return errors.New("cannot withdraw, insufficient funds")
// 	}

// 	w.balance -= amount

// 	return nil
// }

// The var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}
