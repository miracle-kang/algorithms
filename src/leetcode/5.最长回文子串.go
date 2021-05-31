/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	n := len(s)
	start, end := 0, 0
	var expand = func(l, r int) (int, int) {
		for l >= 0 && r < n && s[l] == s[r] {
			l, r = l-1, r+1
		}
		return l + 1, r - 1
	}
	for i := 0; i < n; i++ {
		l1, r1 := expand(i, i)
		l2, r2 := expand(i, i+1)

		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

// @lc code=end

