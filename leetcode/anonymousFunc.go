package leetcode

import "fmt"

func PrintAnonymousFunc()  {

	fmt.Println("\ntest anonymousFunc: ")

	a, b := 10, 20

	go func() {
		fmt.Println("func0 : ", a)
		fmt.Println("func0 : ", b)
	}()

	go func(a, b int) {
		fmt.Println("func1 : ", a)
		fmt.Println("func1 : ", b)
	}(a, b)

	a = 100
	b = 200

	fmt.Println(a)
	fmt.Println(b)


}

func main() {

}
