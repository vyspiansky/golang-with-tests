package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// // The TB name stands for "Test/Benchmark"
	// assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
	// 	t.Helper()
	// 	got := wallet.Balance()

	// 	if got != want {
	// 		t.Errorf("got %s want %s", got, want)
	// 	}
	// }

	// assertError := func(t testing.TB, err error) {
	// 	t.Helper()
	// 	if err == nil {
	// 		t.Error("wanted an error but didn't get one")
	// 	}
	// }

	// assertError := func(t testing.TB, got error, want string) {
	// 	t.Helper()

	// 	if got == nil {
	// 		// t.Fatal will stop the test if it is called.
	// 		t.Fatal("didn't get an error but wanted one")
	// 	}

	// 	// Errors can be converted to a string with the .Error() method
	// 	if got.Error() != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}
	// }

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		// got := wallet.Balance()

		// // Find out the address of that bit of memory via &myVal.
		// // The %p placeholder prints memory addresses in base 16 notation
		// // with leading 0xs and the  escape character prints a new line.
		// fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		// want := Bitcoin(10)

		// if got != want {
		// 	// t.Errorf("got %d want %d", got, want)
		// 	t.Errorf("got %s want %s", got, want)
		// }

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		// wallet.Withdraw(Bitcoin(10))

		// got := wallet.Balance()

		// want := Bitcoin(10)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	// t.Run("withdraw insufficient funds", func(t *testing.T) {
	// 	startingBalance := Bitcoin(20)
	// 	wallet := Wallet{startingBalance}
	// 	err := wallet.Withdraw(Bitcoin(100))

	// 	assertBalance(t, wallet, startingBalance)

	// 	if err == nil {
	// 		t.Error("wanted an error but didn't get one")
	// 	}
	// })

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		// assertError(t, err)
		// assertError(t, err, "cannot withdraw, insufficient funds")
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
