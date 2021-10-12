package bank2_test

import (
	"lgopl/ch9/bank2"
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	var wait sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		wait.Add(1)
		go func(amount int) {
			defer wait.Done()
			bank2.Deposits(amount)
		}(i)
	}
	wait.Wait()

	if got, want := bank2.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	return
}
