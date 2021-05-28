/*
 * @lc app=leetcode.cn id=477 lang=golang
 *
 * [477] 汉明距离总和
 */

// @lc code=start
func totalHammingDistance(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < 32; i++ {
		c := 0
		for _, num := range nums {
			c += (num >> i) & 1
		}
		res += c * (n - c) // 每一位中 1的个数 * 0的个数
	}
	return res
}

// @lc code=end

