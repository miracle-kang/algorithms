/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start
var ans int

func totalNQueens(n int) int {
	ans = 0
	nums := make([]int, n)
	cal(nums, 0)
	return ans
}

func cal(nums []int, row int) {
	n := len(nums)
	if row == n {
		ans++
		return
	}
	for col := 0; col < n; col++ {
		if isOk(nums, row, col) {
			nums[row] = col
			cal(nums, row+1)
		}
	}
}

func isOk(nums []int, row, col int) bool {
	n := len(nums)
	leftup, rightup := col-1, col+1
	for i := row - 1; i >= 0; i-- {
		if nums[i] == col {
			return false
		}
		if leftup >= 0 && nums[i] == leftup {
			return false
		}
		if rightup < n && nums[i] == rightup {
			return false
		}
		leftup--
		rightup++
	}
	return true
}

// @lc code=end

