/*
 * @lc app=leetcode.cn id=879 lang=golang
 *
 * [879] 盈利计划
 */

// @lc code=start
const MOD = 1e9 + 7

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	// dp[i][j][k]=dp[i−1][j][k]+dp[i−1][j−group[i]][max(0,k−profit[i])]
	len := len(group)
	dp := make([][][]int, len+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, minProfit+1)
		}
	}

	dp[0][0][0] = 1
	for i, m := range group {
		p := profit[i]
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				if j < m {
					dp[i+1][j][k] = dp[i][j][k]
				} else {
					dp[i+1][j][k] = (dp[i][j][k] + dp[i][j-m][max(0, k-p)]) % MOD
				}
			}
		}
	}
	sum := 0
	for _, d := range dp[len] {
		sum = (sum + d[minProfit]) % MOD
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

