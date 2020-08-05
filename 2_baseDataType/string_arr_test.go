package __baseDataType

import (
	"fmt"
	"testing"
)

func TestStringIndex(t *testing.T) {
	StringIndex()
}

func TestStringArr(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
