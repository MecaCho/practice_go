package dfs

import (
	"fmt"
	"testing"
)

func TestGetImportance(t *testing.T) {
	res := GetImportance([]*Employee{
		{1, 2, []int{2}},
		{2, 3, []int{}},
	}, 2)
	fmt.Println(res)
}
