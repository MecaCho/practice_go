package backtrack

import "fmt"

func Permute(nums []int) [][]int {

	res := make([][]int, 0)
	path := make([]int, 0)
	helper1(len(nums), nums, path, &res)
	fmt.Println(res)
	return res
}

func helper1(n int, nums, path []int, res *[][]int) {
	// fmt.Println(n, nums, path)
	if len(path) == n {
		*res = append(*res, path)
		return
	}

	for k := range nums {
		path_ := append(path, nums[k])
		new_nums := make([]int, len(nums)-1)
		copy(new_nums[:k], nums[:k])
		copy(new_nums[k:], nums[k+1:])
		helper1(n, new_nums, path_, res)
	}
}
