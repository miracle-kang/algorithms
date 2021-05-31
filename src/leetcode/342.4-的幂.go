/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
const MASK int = 0x55555555
func isPowerOfFour(n int) bool {
	return n&(n-1) == 0 && (n & MASK) != 0
}

// @lc code=end

