
package dp

func canCross(stones []int) bool {
	length := len(stones)
	if stones[1] != 1 {
		return false
	}
	dp := make([][]bool, length)

	for i := range dp {
		dp[i] = make([]bool, length+1)
	}
	dp[0][0] = true

	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k <= j+1 {
				dp[i][k] = dp[j][k] || dp[j][k+1] || dp[j][k-1]
				if i == length-1 && dp[i][k] {
					return true
				}
			}
		}
	}
	return false
}

func CanCross(stones []int) bool {
	return canCross(stones)
}


