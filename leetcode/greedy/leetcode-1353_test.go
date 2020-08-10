package greedy

import (
	"fmt"
	"testing"
)

func TestMaxEvents(t *testing.T) {
	res := MaxEvents([][]int{[]int{1, 2}, []int{2, 3}})
	fmt.Println(res)
}
