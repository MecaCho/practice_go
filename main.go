package main

import (
	"fmt"
	_ "net/http/pprof"
	"time"
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

func MergeArr(arrs [][]int64) []int64 {
	res := []int64{}
	max_len := 0
	for i, _ := range arrs {
		length := len(arrs[i])
		if length > max_len {
			max_len = length
		}
	}

	for i := 0; i < max_len; i++ {
		for _, arr := range arrs {
			if i < len(arr) {
				res = append(res, arr[i])
			}
		}
	}

	fmt.Println(res)
	return res
}

func UsePlug() {

	fmt.Println("test use plugin")

}

func main() {
	// fmt.Println("")
	// // UsePlug()
	// UsePlugin1()
	// UsePlugin2()
	// UsePlugin3()
	// // StartServer()
	// log.Println(http.ListenAndServe("localhost:8081", nil))

	go func() {
		for {
			// time.Sleep(1 * time.Second)
			fmt.Println("hello")
		}
	}()
	time.Sleep(time.Millisecond)
	// runtime.GC()
	println("OK")
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
