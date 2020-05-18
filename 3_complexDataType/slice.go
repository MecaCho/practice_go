package __complexDataType

import "fmt"

func SliceAppend() {

	s := make([]int, 3, 8)
	a := s[:9]
	fmt.Println(s, a)
}

// panic: runtime error: slice bounds out of range [:9] with capacity 8 [recovered]
// 	panic: runtime error: slice bounds out of range [:9] with capacity 8
//
// goroutine 17 [running]:
// testing.tRunner.func1(0xc0000e6100)
// 	/usr/local/go/src/testing/testing.go:874 +0x3a3
// panic(0x11336e0, 0xc0000cc0a0)
// 	/usr/local/go/src/runtime/panic.go:679 +0x1b2
// go_practice/3_complexDataType.SliceAppend()
// 	/Users/rainmc/GO/src/go_practice/3_complexDataType/slice.go:8 +0x4e
// go_practice/3_complexDataType.TestSliceAppend(0xc0000e6100)
// 	/Users/rainmc/GO/src/go_practice/3_complexDataType/slice_test.go:6 +0x20
// testing.tRunner(0xc0000e6100, 0x114f638)
// 	/usr/local/go/src/testing/testing.go:909 +0xc9
// created by testing.(*T).Run
// 	/usr/local/go/src/testing/testing.go:960 +0x350
