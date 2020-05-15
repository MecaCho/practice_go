package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetSliceLen(t *testing.T) {
	GetSliceLen()
}

// func BenchmarkGetSliceLen(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		GetSliceLen()
// 	}
// }

func TestNode_GetInfo(t *testing.T) {

}

func BenchmarkNode_GetInfo(b *testing.B) {
	node := Node{"test_name", 1, [10000]string{}}
	for i := 0; i < b.N; i++ {
		node.GetInfo()
	}
}

func BenchmarkNode_GetInfo1(b *testing.B) {
	node := &Node{"test_name", 1, [10000]string{}}
	for i := 0; i < b.N; i++ {
		node.GetInfo1()
	}
}

const NumOfElems = 1000

type Content struct {
	Detail [10000]int
}

func withValue(arr [NumOfElems]Content) int {
	//	fmt.Println(&arr[2])
	return 0
}

func withReference(arr *[NumOfElems]Content) int {
	//b := *arr
	//	fmt.Println(&arr[2])
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content
	//fmt.Println(&arr[2])
	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
	runtime.GC()
}

func TestMergeArr(t *testing.T) {
	arr1 := []int64{1, 2, 3}
	arr2 := []int64{4, 5}
	arrs := [][]int64{arr1, arr2}
	res := MergeArr(arrs)
	fmt.Println(res)
}
