package __channel

import (
	"fmt"
	"testing"
)

func TestMixPrint(t *testing.T) {
	MixPrint()
}

func TestPrintABC(t *testing.T) {
	res := QuickSort([]int64{1, 4, 5, 712, 1, 3, 4})
	fmt.Println(res)

	arr := &[]int64{1, 2, 3}
	AddItem(arr, 9)
	fmt.Println(arr, *arr)

	PrintABC()

	PrintWaitGroup()
}

func BenchmarkPrintABC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintABC()
	}
}

func BenchmarkPrintWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintWaitGroup()
	}
}

func TestChangeArrItem(t *testing.T) {
	arr := [3]int64{1, 2, 3}
	ChangeArrItem(arr)
	fmt.Println(arr)

	s := []int64{1, 2, 3}
	ChangeSliceItem(s)
	fmt.Println(s)
}
