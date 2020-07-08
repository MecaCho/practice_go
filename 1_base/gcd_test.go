package __base

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	GCD()
}

func TestFib(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := Fib(i)
		fmt.Println(res)
	}

}
