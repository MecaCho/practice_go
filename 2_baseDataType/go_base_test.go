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
	s := make([]int, 10)
	fmt.Println(len(s), cap(s), s)
	for i := 0;i< 10; i++{
		s = append(s, i)
	}
	fmt.Println(s)
}
