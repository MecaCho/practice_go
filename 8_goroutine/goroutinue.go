package goroutine

import (
	"fmt"
)

func main() {
	// fmt.Println("helloworld")
	// go func() {
	// 	for {
	// 		fmt.Println("hello")
	// 	}
	// }()
	//
	// time.Sleep(time.Second)
	fmt.Println("123")
	s := make([]int, 3, 10)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
