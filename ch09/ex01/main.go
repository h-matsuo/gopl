package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

type withdrawResult struct {
	amount int
	ok     bool
}
type withdrawRequest struct {
	amount int
	ch     chan<- *withdrawResult
}

var withdrawal = make(chan *withdrawRequest)

func Withdraw(amount int) bool {
	ch := make(chan *withdrawResult)
	withdrawal <- &withdrawRequest{amount, ch}
	result := <-ch
	return result.ok
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case request := <-withdrawal:
			amount := request.amount
			if balance >= amount {
				balance -= amount
				request.ch <- &withdrawResult{amount, true}
			} else {
				request.ch <- &withdrawResult{0, false}
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
