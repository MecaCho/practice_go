package __baseDataType

import (
	"fmt"
	"testing"
)

func TestDeferTest(t *testing.T) {
	// DeferTest()
	// DeferPanicTest()
	DeferPanicTest1()
}

func TestCapOfSlice(t *testing.T) {
	CapOfSlice()
}

func TestSliceTest(t *testing.T) {
	SliceTest()
}

func TestDeferTest2(t *testing.T) {
	DeferTest()
}

func TestSliceTest2(t *testing.T) {
	s := make([]string, 10)
	fmt.Println(len(s), cap(s), s)
	for i := 0; i < 10; i++ {
		s = append(s, "a")
	}
	fmt.Println(s)
}

func AppendSlice(s []int) {
	s[0] = 100
	for i := 0; i < 11; i++ {
		s = append(s, 100)
	}
	fmt.Printf("slice append, len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)

}

func AppendSlice1(s *[]int) {
	*s = append(*s, 100)
}

func TestSliceTest3(t *testing.T) {
	s := make([]int, 0, 10)
	fmt.Printf("len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)

	s = append(s, 1)
	fmt.Printf("len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)

	AppendSlice(s)
	fmt.Printf("len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)

	AppendSlice1(&s)
	fmt.Printf("len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)
	// 	0 10 []
	// 1 10 [1]
	// 1 10 [100]
	// 2 10 [100 100]
}
