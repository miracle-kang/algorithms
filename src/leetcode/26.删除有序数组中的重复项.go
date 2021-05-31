/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	if len(nums) == 0 {
		return 0
	}
	pre := 0
	for i := 1; i < n; i++ {
		if nums[pre] != nums[i] {
			pre++
			nums[pre] = nums[i]
		}
	}
	return pre+1
}
// @lc code=end

