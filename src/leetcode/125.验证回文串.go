/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	i := 0
	j := len(str) - 1
	for i < j {
		for i < len(str) && !isChar(str[i]) {
			i++
		}
		for j > 0 && !isChar(str[j]) {
			j--
		}

		// 如果是空字符串
		if i >= len(str) || j < 0 {
			return true
		}
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isChar(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c))
}

// @lc code=end

