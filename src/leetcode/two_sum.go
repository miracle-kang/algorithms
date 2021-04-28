package leetcode

// #1.两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	// Dict
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
