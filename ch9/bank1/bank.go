package bank

var (
	deposits = make(chan int)
	balances = make(chan int)
)

func Deposits(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		// 因为管道为无缓冲，所以这里堵塞了，等待Balance调用，会触发消费
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
