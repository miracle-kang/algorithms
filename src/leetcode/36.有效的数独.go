/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	var rmp, cmp, bmp [9][9]bool
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			ch := board[r][c] - '1'
			if rmp[r][ch] || cmp[c][ch] {
				return false
			}
			blk := r - r%3 + c/3
			if bmp[blk][ch] {
				return false
			}
			rmp[r][ch], cmp[c][ch], bmp[blk][ch] = true, true, true
		}
	}
	return true
}

// @lc code=end

