package arr

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{1,2,2,3,4,5,6,6,6,7}
	res := RemoveElement(nums, 2)
	fmt.Println(res)
}
