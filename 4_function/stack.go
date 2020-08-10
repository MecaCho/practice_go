package __function

import "fmt"

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	res := fib(1000)
	fmt.Println(res)
}