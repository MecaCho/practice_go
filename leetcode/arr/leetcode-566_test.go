package arr

import (
	"fmt"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	res := MatrixReshape([][]int{[]int{1,2},[]int{3,4}}, 1,4)
	fmt.Println(res)
}
