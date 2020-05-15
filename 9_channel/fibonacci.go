package __channel

import "fmt"

func Fabonacci() {

	quit := make(chan int, 1)
	ret := make(chan int, 1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ret)
		}
		quit <- 1
	}()
	fibonacci(ret, quit)

}

func fibonacci(ret, quit chan int) {
	a, b := 0, 1
	for {
		select {
		case ret <- a:
			a, b = a+b, a
		case <-quit:
			return
		}
	}
}
