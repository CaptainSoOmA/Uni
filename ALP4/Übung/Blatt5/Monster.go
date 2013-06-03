package main

import "fmt"
import . "sync"
import "strconv" 

func main() {
    fmt.Printf("hello, world\n")
    account := New()
    account.Deposit(30)
    fmt.Printf(strconv.Itoa(account.Balance() ) +"\n")	
}


type Imp struct {
	balance int
	mutex Mutex
	credit *Cond
}

func New() *Imp {
	x:= new(Imp)
	x.credit = NewCond(&x.mutex)
	return x
}

func (x *Imp) Deposit(b int) {
	x.mutex.Lock()
	x.balance += b
	x.credit.Signal()
	x.mutex.Unlock()
}

func (x *Imp) Draw(b int) {
	x.mutex.Lock()
	for x.balance < b {
		x.credit.Wait()
	}
	x.balance -= b
	x.credit.Signal()
	x.mutex.Unlock()
}

func (x *Imp) Balance() int {
	return x.balance
}