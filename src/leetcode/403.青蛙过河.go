/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start
func canCross(stones []int) bool {
	n := len(stones)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == n-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}

// @lc code=end

