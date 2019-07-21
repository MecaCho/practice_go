package leetcode

import "fmt"

var N  = 100

func WithoutBufferChannel()  {

	var ch = make(chan int, 0)

	go func() {
		for i := 0; i < N; i++{
			fmt.Println("write channal without buffer,", i)
			ch <- i
		}
	}()

	for i := 0; i < N; i++{
		fmt.Println("Get channel without buffer: ", <- ch)
	}

}

func BufferChannel()  {

	var ch = make(chan int, 5)

	go func() {
		for i := 0; i < N; i++{
			fmt.Println("write channal with buffer,", i)
			ch <- i
		}
	}()

	for i := 0; i < N; i++{
		fmt.Println("Get channel with buffer: ", <- ch)
	}

}
