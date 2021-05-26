/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	stack := []string{}
	i, j := 0, 0

	for i < len(s) {
		if unicode.IsDigit(rune(s[i])) {
			j = i
			i++

			// 可能出现多个数字的情况，此时获取全部数字并压栈
			for unicode.IsDigit(rune(s[i])) {
				i++
			}
			stack = append(stack, string(s[j:i]))
		} else if s[i] == '[' || unicode.IsLetter(rune(s[i])) {
			// 字母和 '[' 直接压栈
			stack = append(stack, string(s[i]))
			i++
		} else {
			str := ""
			// 如果是 ']'，从栈中取元素，拼接成字符串，直到遇到了 '['
			for len(stack) > 0 && stack[len(stack)-1] != "[" {
				str = stack[len(stack)-1] + str
				stack = stack[0 : len(stack)-1]
			}
			stack = stack[0 : len(stack)-1]

			// 前面取完了 '[' 之后的所有字符，此时栈中第一个就是数字
			rep, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[0 : len(stack)-1]

			str = strings.Repeat(str, rep)
			stack = append(stack, str)
			i++
		}
	}

	result := ""
	for i = 0; i < len(stack); i++ {
		result += stack[i]
	}
	return result
}

// @lc code=end

