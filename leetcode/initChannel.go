package leetcode

import "fmt"

func InitChannel() {
	var ch chan int
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
	}()

	fmt.Println("read channel without buffer : ", <-ch1)

	//go func() {
	//	ch <- 1
	//}()

	//<- ch
	fmt.Println("init channel: ", ch)
	//ch <-
	//close(ch)
	//close(ch1)
}
