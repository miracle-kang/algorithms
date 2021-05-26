package leetcode

// 1190. 反转每对括号间的子串
// 给出一个字符串 s（仅含有小写英文字母和括号）。
// 请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
// 注意，您的结果中 不应 包含任何括号。
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
