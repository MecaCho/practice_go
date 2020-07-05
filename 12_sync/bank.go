package _2_sync

import (
	"fmt"
	"sync"
	"time"
)

var (
	sema    = make(chan struct{}, 1)
	mu     sync.Mutex
	balance int
)

func Deposit0(amount int) {
	balance += amount
}

func Balance0() int {
	return balance
}

func Deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

func Deposit1(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance1() int {
	mu.Lock()
	defer mu.Unlock()

	return balance
}

func RunBank() {
	balance = 10000
	for i := 0; i < 5; i++ {
		go func() {
			Deposit0(100)
			fmt.Println("+100, result: ", Balance0())
		}()

		go func() {
			Deposit0(-50)
			fmt.Println("-50, result: ", Balance0())
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println(balance)

}
