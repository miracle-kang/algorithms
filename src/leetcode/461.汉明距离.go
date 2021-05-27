/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	z := x ^ y
	cnt := 0
	for z > 0 {
		if z & 1 == 1 {
			cnt++
		}
		z = z >> 1
	}
	return cnt
}
// @lc code=end

