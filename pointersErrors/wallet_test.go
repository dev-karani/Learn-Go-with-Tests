package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Test wallet balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		fmt.Printf("address of balance intest is %p \n", &wallet.balance)

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)

		}

	})
	t.Run("Bitcoin String", func(t *testing.T) {
		btc := Bitcoin(10)
		got := btc.String()
		want := "20 BTC"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
