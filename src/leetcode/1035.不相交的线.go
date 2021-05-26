/*
 * @lc app=leetcode.cn id=1035 lang=golang
 *
 * [1035] 不相交的线
 */

// @lc code=start
func maxUncrossedLines(nums1 []int, nums2 []int) int {
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

// @lc code=end

