package link_list

func Rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	dp := [length]int{}
	// dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = nums[1]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	}
	for j := 2; j < length; j++ {
		doIt := dp[j-2] + nums[j]
		notDo := dp[j-1]
		if doIt > notDo {
			dp[j] = doIt
		} else {
			dp[j] = notDo
		}
	}
	return dp[length-1]
}

func rob1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	// dp := make([]int, length)
	// dp[0] = nums[0]
	// dp[1] = nums[1]
	pre := nums[0]
	cur := nums[1]
	if nums[0] > nums[1] {
		// dp[1] = nums[0]
		cur = nums[0]
	}
	for j := 2; j < length; j++ {
		// doIt := dp[j-2] + nums[j]
		// notDo := dp[j-1]
		doIt := pre + nums[j]
		notDo := cur
		pre = cur
		if doIt > notDo {
			// dp[j] = doIt
			cur = doIt
		} else {
			// dp[j] = notDo
			cur = notDo
		}
	}
	// return dp[length-1]
	return cur

}
