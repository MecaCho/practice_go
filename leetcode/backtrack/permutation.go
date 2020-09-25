package backtrack

import "fmt"

// func permute(nums []int) [][]int {
//
// 	return Impl(3)
// }

func helper(n int, path *[]int, rst *[][]int) {
	if n == len(*path) {
		*rst = append(*rst, *path)
		return
	}
	//
	// length := len(*path)
	for i := 1; i < n+1; i++ {
		path_ := append(*path, i)
		helper(n, &path_, rst)
		// *path = (*path)[:len(*path)-1]
	}
}

func Impl(n int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)

	helper(n, &path, &ret)
	fmt.Println(ret)
	return ret
}
