package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(Bitcoin(100))

	wallet.Withdraw(Bitcoin(10))

	got := wallet.Balance()
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	want := Bitcoin(90)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
