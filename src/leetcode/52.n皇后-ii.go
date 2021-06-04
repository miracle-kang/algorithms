/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start
func totalNQueens(n int) int {
	ans := 0
	cmp := make([]bool, n)
	lmp, omp := make([]bool, n*2), make([]bool, n*2)

	var cal func(int)
	cal = func(row int) {
		if row == n {
			ans++
			return
		}
		for col := 0; col < n; col++ {
			if cmp[col] {
				continue
			}
			ll, ol := row+col, n-row+col-1
			if lmp[ll] || omp[ol] {
				continue
			}
			cmp[col], lmp[ll], omp[ol] = true, true, true
			cal(row + 1)
			cmp[col], lmp[ll], omp[ol] = false, false, false
		}
	}
	cal(0)
	return ans
}

// @lc code=end

