/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	m := map[int]int{0: -1}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		rem := sum % k
		if pre, has := m[rem]; has {
			if i - pre > 1 {
				return true
			}
		} else {
			m[rem] = i
		}
	}
	return false
}

// @lc code=end

func checkSubarraySum1(nums []int, k int) bool {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		sum := nums[i]
		for j := i + 1; j < n; j++ {
			sum += nums[j]
			if sum%k == 0 {
				return true
			}
		}
	}
	return false
}