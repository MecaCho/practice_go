package leetcode

import (
	"time"
	"fmt"
)

func Produce1(ch chan int)  {

	go func() {
		for i := 0; i < 10; i++{
			time.After(1 * time.Second)
			ch <- i
			fmt.Println("Produce1 channel, ---->", i)
		}
	}()

}

func Produce2(ch chan int, done chan bool)  {

	i := 0
	go func(int) {
		for {
			time.Sleep(1 * time.Second)
			ch <- i
			i++
			fmt.Println("Produce2 channel, ---->", i)
			if i > 100{
				done <- true
				break
			}
		}
	}(i)

}

func Consumer(ch chan int, done chan bool){

	for{
		select {
		case v, ok := <- ch:
			fmt.Println("consume channel : ---------------", v, ok)
		case <- done:
			return
		//default:
		//	fmt.Println("channel not found.", time.Now())
			//break
		}
	}



}

func Run() {

	ch := make(chan int, 10)
	done := make(chan bool)

	Produce1(ch)
	Produce2(ch, done)
	Consumer(ch, done)

}