/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
var rmp, cmp, bmp [9][9]bool

func solveSudoku(board [][]byte) {
	rmp, cmp, bmp = [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			ch := board[r][c] - '1'
			blk := r - r%3 + c/3
			rmp[r][ch], cmp[c][ch], bmp[blk][ch] = true, true, true
		}
	}
	sudoku(board, 0)
}

func sudoku(board [][]byte, n int) bool {
	row, col := n/9, n%9
	blk := row - row%3 + col/3
	if row > 8 {
		return true
	}
	if board[row][col] != '.' {
		return sudoku(board, n+1)
	}
	for c := '1'; c <= '9'; c++ {
		ch := c - '1'
		if rmp[row][ch] || cmp[col][ch] || bmp[blk][ch] {
			continue
		}
		board[row][col] = byte(c)
		rmp[row][ch], cmp[col][ch], bmp[blk][ch] = true, true, true
		if done := sudoku(board, n+1); done {
			return done
		}
		board[row][col] = '.'
		rmp[row][ch], cmp[col][ch], bmp[blk][ch] = false, false, false
	}
	return false
}

// @lc code=end
