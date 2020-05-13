package main

import (
	"fmt"
	"go_practice/gogc"
)

type Node struct {
	Name string
	ID   int64
	Info [10000]string
}

func (n Node) GetInfo() string {

	return fmt.Sprintf("%s, %d", n.Name, n.ID)
}

func (n *Node) GetInfo1() string {

	return fmt.Sprintf("%s, %d", n.Name, n.ID)
}

func GetSliceLen() {

	s := []int{}
	for i := 0; i < 10000; i++ {
		s = append(s, i)
		// fmt.Println(i, len(s), cap(s))
	}
	fmt.Println(len(s), cap(s))
}

// func fibonacci(ret, quit chan int) {
// 	a, b := 0, 1
// 	for {
// 		select {
// 		case ret <- a:
// 			a, b = a+b, a
// 		case <-quit:
// 			return
// 		}
// 	}
// }

func main() {
	gogc.GetBonus()
	// quit := make(chan int, 1)
	// ret := make(chan int, 1)
	//
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-ret)
	// 	}
	// 	quit <- 1
	// }()
	// fibonacci(ret, quit)
}

//func main() {
//	quit := make(chan int, 1)
//	ret := make(chan int, 1)
//
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(<-ret)
//		}
//		quit <- 1
//	}()
//	fibonacci(ret, quit)
//	var qwq error = nil
//}
