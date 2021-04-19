package arr

import "fmt"

func removeElement(nums []int, val int) int {

	i, n := 0, len(nums)

	for ; i < n; {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n -= 1
		} else {
			i += 1
		}
	}
	fmt.Println(nums)
	return n

}

func RemoveElement(nums []int, val int) int {
	return removeElement(nums, val)
}
