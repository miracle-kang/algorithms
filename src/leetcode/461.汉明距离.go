/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	z := x ^ y
	cnt := 0
	for ; z > 0; z&=z-1 {
		cnt++
	}
	return cnt
}

// @lc code=end

func hammingDistance1(x int, y int) int {
	z := x ^ y
	cnt := 0
	for ; z > 0; z>>=1 {
		cnt += z&1
	}
	return cnt
}