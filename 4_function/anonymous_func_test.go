package __function

import "testing"

func TestAnonymFunc(t *testing.T) {
	AnonymFunc()

	AnonymFunc1()
}

func TestMyFunc(t *testing.T) {
	args := []int{1, 2, 3}
	MyFunc(args...)

	MyFunc(4, 5)
	MyFunc(6, 7, 8, 9)
}
