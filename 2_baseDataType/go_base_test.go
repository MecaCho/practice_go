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

func TestAppendSlice(t *testing.T) {
	AppendSlice20210122()
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
		fmt.Printf("slice append, len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)
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

func TestSliceTest4(t *testing.T) {
	s := make([]int, 0, 10)
	for i := 0; i < 11; i++ {
		s = append(s, i)
		fmt.Printf("slice append, len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)
	}
	fmt.Printf("slice append, len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)
	// 	slice append, len: 1, cap: 10, value: [0], address: 0xc0000f6000.
	// slice append, len: 2, cap: 10, value: [0 1], address: 0xc0000f6000.
	// slice append, len: 3, cap: 10, value: [0 1 2], address: 0xc0000f6000.
	// slice append, len: 4, cap: 10, value: [0 1 2 3], address: 0xc0000f6000.
	// slice append, len: 5, cap: 10, value: [0 1 2 3 4], address: 0xc0000f6000.
	// slice append, len: 6, cap: 10, value: [0 1 2 3 4 5], address: 0xc0000f6000.
	// slice append, len: 7, cap: 10, value: [0 1 2 3 4 5 6], address: 0xc0000f6000.
	// slice append, len: 8, cap: 10, value: [0 1 2 3 4 5 6 7], address: 0xc0000f6000.
	// slice append, len: 9, cap: 10, value: [0 1 2 3 4 5 6 7 8], address: 0xc0000f6000.
	// slice append, len: 10, cap: 10, value: [0 1 2 3 4 5 6 7 8 9], address: 0xc0000f6000.
	// slice append, len: 11, cap: 20, value: [0 1 2 3 4 5 6 7 8 9 10], address: 0xc0000fe000.
	// slice append, len: 11, cap: 20, value: [0 1 2 3 4 5 6 7 8 9 10], address: 0xc0000fe000.
}

func TestSliceTest5(t *testing.T) {
	s := []int{}
	fmt.Printf("slice init, len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)
	// slice init, len: 0, cap: 0, value: [], address: 0x1264148.

	var s1 []int
	fmt.Printf("slice init, len: %d, cap: %d, value: %+v, address: %p.\n", len(s1), cap(s1), s1, s1)
	// slice init, len: 0, cap: 0, value: [], address: 0x0.

	s3 := make([]int, 10)
	fmt.Printf("slice init, len: %d, cap: %d, value: %+v, address: %p.\n", len(s3), cap(s3), s3, s3)
	// slice init, len: 10, cap: 10, value: [0 0 0 0 0 0 0 0 0 0], address: 0xc0000fc000.

	s4 := make([]int, 1, 10)
	fmt.Printf("slice init, len: %d, cap: %d, value: %+v, address: %p.\n", len(s4), cap(s4), s4, s4)
	// slice init, len: 1, cap: 10, value: [0], address: 0xc0000ee140.

	s5 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice init, len: %d, cap: %d, value: %+v, address: %p.\n", len(s5), cap(s5), s5, s5)
	// slice init, len: 5, cap: 5, value: [1 2 3 4 5], address: 0xc0000fe000.
}

func TestSliceAppend2(t *testing.T) {

	s := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Printf("slice s, len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)

	s1 := s[2:4]
	fmt.Printf("slice s1, len: %d, cap: %d, value: %+v, address: %p.\n", len(s1), cap(s1), s1, s1)

	s2 := s[2:5]
	fmt.Printf("slice s2, len: %d, cap: %d, value: %+v, address: %p.\n", len(s2), cap(s2), s2, s2)

	ss1 := s1[0:5]
	fmt.Printf("slice ss1, len: %d, cap: %d, value: %+v, address: %p.\n", len(ss1), cap(ss1), ss1, ss1)

	ss1 = append(ss1, 100)
	fmt.Printf("slice ss1, len: %d, cap: %d, value: %+v, address: %p.\n", len(ss1), cap(ss1), ss1, ss1)
	fmt.Printf("slice s, len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)

	s2 = append(s2, 101)
	fmt.Printf("slice s2, len: %d, cap: %d, value: %+v, address: %p.\n", len(s2), cap(s2), s2, s2)

	s2 = append(s2, 100)
	fmt.Printf("slice s2, len: %d, cap: %d, value: %+v, address: %p.\n", len(s2), cap(s2), s2, s2)

	s2 = append(s2, 100)
	fmt.Printf("slice s2, len: %d, cap: %d, value: %+v, address: %p.\n", len(s2), cap(s2), s2, s2)

	s2 = append(s2, 100)
	fmt.Printf("slice s2, len: %d, cap: %d, value: %+v, address: %p.\n", len(s2), cap(s2), s2, s2)

	fmt.Printf("slice s, len: %d, cap: %d, value: %+v, address: %p.\n", len(s), cap(s), s, s)
	// 	slice s, len: 7, cap: 7, value: [0 1 2 3 4 5 6], address: 0xc000106000.
	// slice s1, len: 2, cap: 5, value: [2 3], address: 0xc000106010.
	// slice s2, len: 3, cap: 5, value: [2 3 4], address: 0xc000106010.
	// slice ss1, len: 5, cap: 5, value: [2 3 4 5 6], address: 0xc000106010.
	// slice ss1, len: 6, cap: 10, value: [2 3 4 5 6 100], address: 0xc00010c000.
	// slice s, len: 7, cap: 7, value: [0 1 2 3 4 5 6], address: 0xc000106000.
	// slice s2, len: 4, cap: 5, value: [2 3 4 101], address: 0xc000106010.
	// slice s2, len: 5, cap: 5, value: [2 3 4 101 100], address: 0xc000106010.
	// slice s2, len: 6, cap: 10, value: [2 3 4 101 100 100], address: 0xc00010c050.
	// slice s2, len: 7, cap: 10, value: [2 3 4 101 100 100 100], address: 0xc00010c050.
	// slice s, len: 7, cap: 7, value: [0 1 2 3 4 101 100], address: 0xc000106000.
}
