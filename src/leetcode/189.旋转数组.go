/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	x := k%len(nums)
	reverse(nums)
	reverse(nums[0:x])
	reverse(nums[x:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

// @lc code=end

func rotate1(nums []int, k int) {
	n := len(nums)
	tmp := make([]int, n)
	for i := range nums {
		tmp[(i+k)%n] = nums[i]
	}
	copy(nums, tmp)
}