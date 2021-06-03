/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	area := 0
	left, right := 0, len(height)-1

	for left < right {
		a := min(height[left], height[right]) * (right - left)
		if a > area {
			area = a
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

