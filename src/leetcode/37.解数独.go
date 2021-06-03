/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	var sudoku func(row, col int) bool
	next := func(row, col int) bool {
		if col >= 8 {
			return sudoku(row+1, 0)
		} else {
			return sudoku(row, col+1)
		}
	}
	sudoku = func(row, col int) bool {
		if row > 8 {
			return true
		}
		if board[row][col] != '.' {
			if done := next(row, col); done {
				return done
			}
			return false
		}
		for c := '1'; c <= '9'; c++ {
			ch := byte(c)
			if isOk(board, row, col, ch) {
				board[row][col] = ch
				if done := next(row, col); done {
					return done
				}
				board[row][col] = '.'
			}
		}
		return false
	}
	sudoku(0, 0)
}

func isOk(board [][]byte, row, col int, ch byte) bool {
	br, bc := row-row%3, col-col%3
	for i := 0; i < 9; i++ {
		if board[row][i] == ch {
			return false
		}
		if board[i][col] == ch {
			return false
		}
		if board[br+i/3][bc+i%3] == ch {
			return false
		}
	}
	return true
}

// @lc code=end
