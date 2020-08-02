package __helloworld

import "fmt"

func MakeNew() {

	slices := make([]int64, 10)
	for k := range slices {
		fmt.Println(k, slices[k])
	}

	maps := make(map[int64]int64, 10)
	for k, v := range maps {
		fmt.Println(k, v)
	}

	// chans := make(chan int64, 10)
	// for k := range chans {
	// 	fmt.Println(k)
	// }

	slices1 := new([]int64)
	fmt.Println(slices1)

	a := new(int64)
	fmt.Println(a, *a)
}
