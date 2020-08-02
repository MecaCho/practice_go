package __helloworld

import (
	"fmt"
	"unsafe"
)

func Arr() {
	c := make(chan int)
	close(c)
	close(c)
	arr := [3]int{1, 2, 3}
	arr0 := [...]int{1, 2, 3}

	fmt.Println(arr, arr0)
}

func Matrix() {
	fmt.Println("Size of int: ", unsafe.Sizeof(int(0)))
	matirx := [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}

	for i, v := range matirx {

		for j, vv := range v {
			fmt.Println(vv)
			fmt.Println(i, j, matirx[i][j], &matirx[i][j])
		}
	}

}
