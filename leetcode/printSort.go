package leetcode

import (
	"fmt"
	"sort"
)

type SSort struct {
	v int
}

type ByAge []SSort

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].v < a[j].v }

func PrintSort() {
	s := []SSort{{1}, {3}, {5}, {2}}

	fmt.Printf("%#v", s)
	// A
	sort.Sort(ByAge(s))
	fmt.Println("")
	fmt.Printf("%#v", s)
	fmt.Println("")
}
