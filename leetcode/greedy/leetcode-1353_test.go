package greedy

import (
	"testing"
	"fmt"
)

func TestMaxEvents(t *testing.T) {
	res := MaxEvents([][]int{[]int{1, 2}, []int{2, 3}})
	fmt.Println(res)
}


