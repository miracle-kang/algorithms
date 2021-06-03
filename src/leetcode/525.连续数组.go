/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 */

// @lc code=start
func findMaxLength(nums []int) int {
	cnt := 0
	maxLen := 0
	mp := map[int]int{0: -1}
	for i, num := range nums {
		if num == 1 {
			cnt++
		} else {
			cnt--
		}

		if pre, has := mp[cnt]; has {
			if n := i - pre; n > maxLen {
				maxLen = n
			}
		} else {
			mp[cnt] = i
		}
	}
	return maxLen
}

// @lc code=end

