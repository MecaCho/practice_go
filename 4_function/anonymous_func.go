package __function

import (
	"fmt"
	"math"
)

func AnonymFunc() {
	j := 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}

func AnonymFunc1() {
	j := 5
	a := func() {
		i := 10
		fmt.Printf("i: %d, j: %d\n", i, j)
	}
	a()
	j *= 2
	a()
}

const p = 0.0000001

func IsEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < p
}

func MyFunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
