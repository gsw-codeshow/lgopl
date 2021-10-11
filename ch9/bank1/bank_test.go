package bank_test

import (
	bank "lgopl/ch9/bank1"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	go func() {
		bank.Deposits(100)
		println("=", bank.Balance())
		done <- struct{}{}
	}()
	go func() {
		bank.Deposits(200)
		done <- struct{}{}
	}()
	<-done
	<-done
	if got, want := bank.Balance(), 300; got != want {
		t.Fatal("bank failed!")
	}
	return
}
