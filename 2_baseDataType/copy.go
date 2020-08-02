package __baseDataType

import (
	"fmt"
	"runtime"
)

func CopyAndDeepCopy() {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	a = b

	fmt.Println(a, b, check)
	// [3, 4], [3, 4], [1, 2]
	example2()
}

func example2() {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	copy(a, b)

	// fmt.Println(&a, &b, &check)
	fmt.Println(a, b, check)
	// [3, 4], [3, 4], [3, 4]
}

func CopySlice() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	copy(a, b)
	fmt.Println(a, b)
	a[0] = 100
	fmt.Println(a, b)
}

func CopySlice1() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	b = a
	fmt.Println(a, b)
	a[0] = 100
	fmt.Println(a, b)
	runtime.NumGoroutine()
}
