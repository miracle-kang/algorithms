/*
 * @lc app=leetcode.cn id=1269 lang=golang
 *
 * [1269] 停在原地的方案数
 */

// @lc code=start
func numWays(steps int, arrLen int) int {
	const mod = 1e9 + 7
	maxCol := min(steps, arrLen-1)
	dp := make([]int, maxCol+1)
	dp[0] = 1
	for i := 1; i <= steps; i++ {
		dpNext := make([]int, maxCol+1)
		for j := 0; j <= maxCol; j++ {
			dpNext[j] = dp[j]
			if j > 0 {
				dpNext[j] = (dpNext[j] + dp[j-1]) % mod
			}
			if j < maxCol {
				dpNext[j] = (dpNext[j] + dp[j+1]) % mod
			}
		}
		dp = dpNext
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

