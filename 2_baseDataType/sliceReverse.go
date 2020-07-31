package __baseDataType

import "fmt"

func reverse(a []int) {
	length := len(a)
	for i := 0; i < length/2; i++ {
		tmp := a[i]
		a[i] = a[length-i-1]
		a[length-i-1] = tmp
	}
}

func Reverse() {
	a := []int{1, 2, 3, 4, 5, 6}
	reverse(a)
	fmt.Println(a)
	// 	// Output: [3 2 1]
}
