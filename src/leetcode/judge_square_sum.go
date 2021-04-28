package leetcode

import "math"

// #633. 平方数之和
// 给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。
func judgeSquareSum(c int) bool {
	if c <= 2 {
		return true
	}

	i := 0
	j := int(math.Sqrt(float64(c)))
	for i <= j {
		sum := i*i + j*j
		if sum < c {
			i++
		} else if sum > c {
			j--
		} else {
			return true
		}
	}
	return false
}
