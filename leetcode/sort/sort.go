package sort

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// // Sort is a convenience method.
// func (p IntSlice) Sort() { Sort(p) }

func Sort()  {
	arr := []int{1,9,1}
	sort.Ints(arr)



	sort.Sort(IntSlice(arr))
	fmt.Println(arr)
}
