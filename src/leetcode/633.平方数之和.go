/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
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

// @lc code=end

