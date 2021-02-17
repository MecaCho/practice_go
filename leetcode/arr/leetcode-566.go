package arr

import "fmt"

func MatrixReshape(nums [][]int, r int, c int) [][]int {
	return matrixReshape(nums, r, c)
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}

	r0, c0 := len(nums), len(nums[0])
	if r0*c0 != r*c {
		return nums
	}

	res := make([][]int, r)
	for i, num := range nums {
		for j, nu := range num {
			index := (i*c0 + j) / c
			fmt.Println(index, nu, i*c0+j, c)
			res[index] = append(res[index], nu)
		}
	}
	return res
}
