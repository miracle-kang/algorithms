/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
var ans [][]string

func solveNQueens(n int) [][]string {
	nums := make([]int, n)
	ans = make([][]string, 0)
	cal(nums, 0)
	return ans
}

func cal(nums []int, row int) {
	n := len(nums)
	if row == n {
		printQueue(nums)
		return
	}
	for col := 0; col < n; col++ {
		if isOk(nums, row, col) {
			nums[row] = col
			cal(nums, row+1)
		}
	}
}

func printQueue(nums []int) {
	n := len(nums)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		line := make([]rune, n)
		for j := 0; j < n; j++ {
			if nums[i] == j {
				line[j] = 'Q'
			} else {
				line[j] = '.'
			}
		}
		res[i] = string(line)
	}
	ans = append(ans, res)
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

