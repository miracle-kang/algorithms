/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
var T = [4]string{"", "M", "MM", "MMM"}
var H = [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var D = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var N = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

func intToRoman(num int) string {
	return T[num/1000] + H[num%1000/100] + D[num%100/10] + N[num%10]
}

// @lc code=end

var nums = [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romans = [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman1(num int) string {
	res := ""
	for i, n := range nums {
		cnt := num / n
		for j := 0; j < cnt; j++ {
			res += romans[i]
		}
		num -= cnt * n
	}
	return res
}
