package __base

import "fmt"

func Arr() {

	arr := [3]int{1, 2, 3}
	arr0 := [...]int{1, 2, 3}

	fmt.Println(arr, arr0)
}
