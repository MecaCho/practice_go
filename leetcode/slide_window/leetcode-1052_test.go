package slide_window

import (
	"fmt"
	"testing"
)

func TestCheckInclusion2(t *testing.T) {
	res := MaxSatisfied([]int{1,0,1,2,1,1,7,5}, []int{0,1,0,1,0,1,0,1}, 3)
	fmt.Println(res)
}
