package __helloworld

import "fmt"

func GCD() {

	gcd := func(x, y int) int {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}
	res := gcd(10, 5)
	fmt.Println(res)
}

func Fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
