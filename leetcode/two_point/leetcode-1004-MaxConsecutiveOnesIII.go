package two_point

func longestOnes(A []int, K int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	zeroIndexs := []int{}
	maxLength := 1

	i := -1
	lastZeroIndex := -1
	preLastZeroIndex := -1

	length := len(nums)

	for k := 0; k < length; k++ {
		if nums[k] == 1 {
			newLength := k - i
			if lastZeroIndex == i {
				newLength++
			}
			if newLength > maxLength {
				maxLength = newLength
			}
		} else {
			if len(zeroIndexs) < K {
				zeroIndexs = append(zeroIndexs, k)
			} else {
				zeroIndexs = zeroIndexs[1:K]
				zeroIndexs = append(zeroIndexs, k)
			}

			preLastZeroIndex = zeroIndexs[0]
			lastZeroIndex = k
			if k <= K || nums[k-K] == 1 {
				i = preLastZeroIndex
			} else {
				i = lastZeroIndex
			}
		}
	}
	if lastZeroIndex == -1 {
		maxLength = length
	}
	return maxLength

}
