package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	withdraws <- amount
	if <-withdraws < 0 {
		return false
	}
	return true
}

func teller() {
	var balance int
	for true {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case x := <-withdraws:
			if x > balance {
				withdraws <- -1
			} else {
				withdraws <- 0
			}
		}
	}
}

func init() {
	go teller()
}
