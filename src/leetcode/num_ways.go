package leetcode

// 1269. 停在原地的方案数
// 有一个长度为 arrLen 的数组，开始有一个指针在索引 0 处。
// 每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。
// 给你两个整数 steps 和 arrLen ，请你计算并返回：在恰好执行 steps 次操作以后，指针仍然指向索引 0 处的方案数。
// 解题思路，当前步的方案数=上一步不动的方案数+上一步往左移的方案数+上一步往右移的方案数，类似于爬楼梯
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
