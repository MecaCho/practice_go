package __channel

import (
	"fmt"
	"runtime"
	"time"
)

// === RUN   TestCloseChan
// panic: close of nil channel
func CloseChan() {

	var ch chan int
	var count int

	go func() {
		ch <- 1
	}()

	go func() {
		count++
		close(ch)
	}()

	<-ch
	fmt.Println(count)

}

func ChanRun() {

	// var ch chan int

	ch := make(chan int)
	go func() {
		for {
			// ch = make(chan int, 1)
			ch <- 1
			fmt.Println("write chan", ch)
			time.Sleep(2 * time.Second)
		}

	}()

	go func() {
		for {
			select {
			case ret := <-ch:

				time.Sleep(time.Second)

				fmt.Println("read chan", ret)
			}
		}

	}()

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}

}
