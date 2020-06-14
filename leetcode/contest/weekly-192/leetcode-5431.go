package weekly_192

import "math"

// class Solution(object):
//    def minCost(self, houses, cost, m, n, target):
//        """
//        :type houses: List[int]
//        :type cost: List[List[int]]
//        :type m: int
//        :type n: int
//        :type target: int
//        :rtype: int
//        """
//        if len(set(houses)) > target:
//            return -1
//        dp = [[[float("inf") for _ in range(target)] for _ in range(n)] for _ in range(m)]
//        for k in range(n):
//            if houses[0] == 0:
//                dp[0][k][0] = cost[0][k]
//        if houses[0] != 0:
//            dp[0][houses[0]-1][0] = 0
//
//        for i in range(1, m):
//            for j in range(n):
//                for k in range(target):
//                    if houses[i] != 0 and houses[i] - 1 != j:
//                        break
//                    tmp = []
//                    for kk in range(n):
//                        if kk != j and k - 1 < target and k > 0:
//                            tmp.append(dp[i-1][kk][k-1])
//                    tmp.append(dp[i-1][j][k])
//                    cur_cost = cost[i][j]
//                    if houses[i] != 0:
//                        cur_cost = 0
//                    dp[i][j][k] = min(tmp) + cur_cost
//
//        res = float("inf")
//        for i in range(n):
//            res = min(res, dp[m-1][i][target-1])
//        return res if res != float("inf") else -1

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := [][][]int{}
	for i := 0; i < m; i++ {
		dp = append(dp, [][]int{})
		for j := 0; j < n; j++ {

			dp[i] = append(dp[i], []int{})

			for k := 0; k < target; k++ {
				dp[i][j] = append(dp[i][j], math.MaxInt64)
			}
		}
	}

	for j := 0; j < n; j++ {
		if houses[0] == 0 {
			dp[0][j][0] = cost[0][j]
		} else {
			dp[0][houses[0]-1][0] = 0
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < target; k++ {
				if houses[i] != 0 && houses[i]-1 != j {
					break
				}
				minVal := dp[i-1][j][k]
				for kk := 0; kk < n; kk++ {
					if kk != j && k > 0 && k-1 < target {
						if dp[i-1][kk][k-1] < minVal {
							minVal = dp[i-1][kk][k-1]
						}
						// tmp = append(tmp, dp[i-1][kk][k-1])
					}
				}
				curCost := cost[i][j]
				if houses[i] != 0 {
					curCost = 0
				}
				if minVal != math.MaxInt64 {
					dp[i][j][k] = minVal + curCost
				}

			}
		}
	}

	res := -1
	for i := 0; i < n; i++ {
		if res == -1 || dp[m-1][i][target-1] < res {
			res = dp[m-1][i][target-1]
		}
	}

	if res >= math.MaxInt64 {
		return -1
	}

	return res
}
