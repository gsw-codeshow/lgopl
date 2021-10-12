package bank2

var (
	semaphore = make(chan struct{}, 1)
	balance   int
)

func Deposits(amount int) {
	semaphore <- struct{}{}
	balance += amount
	<-semaphore
}

func Balance() int {
	semaphore <- struct{}{}
	b := balance
	<-semaphore
	return b
}
