package Solutions

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	r := make([][]byte, 9)
	l := make([][]byte, 9)
	b := make([][]byte, 9)
	bIndex := 0
	var tem byte
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			for _, tem = range r[i] {
				if tem == board[i][j] {
					return false
				}
			}
			r[i] = append(r[i], board[i][j])
			for _, tem = range l[j] {
				if tem == board[i][j] {
					return false
				}
			}
			l[j] = append(l[j], board[i][j])
			bIndex = (i/3)*3 + j/3
			for _, tem = range b[bIndex] {
				if tem == board[i][j] {
					return false
				}
			}
			b[bIndex] = append(b[bIndex], board[i][j])
		}
	}
	return true
}

// @lc code=end
