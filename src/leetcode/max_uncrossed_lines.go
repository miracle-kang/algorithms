package leetcode

// 1035. 不相交的线
// 在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
// 现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：
// 		nums1[i] == nums2[j]
//		且绘制的直线不与任何其他连线（非水平线）相交。
// 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
// 以这种方法绘制线条，并返回可以绘制的最大连线数。
//
// PS: 其实就是求最长公共子序列(1143)
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[n][m]
}

// 压缩数组优化
func maxUncrossedLines1(nums1 []int, nums2 []int) int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m := len(nums2)
	dp := make([]int, m+1)
	next := make([]int, m+1)
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				next[j+1] = dp[j] + 1
			} else {
				next[j+1] = max(dp[j+1], next[j])
			}
		}
		dp, next = next, dp
	}
	return dp[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
