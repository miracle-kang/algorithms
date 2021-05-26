/*
 * @lc app=leetcode.cn id=1190 lang=golang
 *
 * [1190] 反转每对括号间的子串
 */

// @lc code=start
func reverseParentheses(s string) string {
	res, _ := reverse(s, 0)
	return string(res)
}

func reverse(s string, start int) ([]byte, int) {
	res := []byte{}
	for i := start; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			str, idx := reverse(s, i+1)
			res = append(res, str...)
			i = idx
		} else if c == ')' {
			for j, n := 0, len(res); j < n/2; j++ {
				res[j], res[n-j-1] = res[n-j-1], res[j]
			}
			return res, i
		} else {
			res = append(res, c)
		}
	}
	return res, len(s) - 1
}

// @lc code=end

