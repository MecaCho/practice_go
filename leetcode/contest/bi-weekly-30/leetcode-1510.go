package bi_weekly_30

func winnerSquareGame(n int) bool {
	dp := []int{0}

	for i := 1; i <= n; i++ {
		dp = append(dp, 0)
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j] == 0 {
				dp[i] = 1
				break
			}
		}
	}
	return dp[n] == 1
}
