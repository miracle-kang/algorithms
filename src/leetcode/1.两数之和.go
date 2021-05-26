/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	i := 0
	for i < len(nums) {
		pre, ok := dict[nums[i]]
		if ok {
			return []int{pre, i}
		}
		dict[target-nums[i]] = i
		i++
	}
	return []int{}
}

// @lc code=end

func twoSum1(nums []int, target int) []int {
	// Double while
	i := 0
	for i < len(nums)-1 {
		j := i + 1
		for j < len(nums) {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
			j++
		}
		i++
	}

	return []int{}
}
