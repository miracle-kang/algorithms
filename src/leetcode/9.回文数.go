/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

import "strconv"

// @lc code=start
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i, n := 0, len(s); i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

// @lc code=end

