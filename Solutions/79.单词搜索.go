package Solutions

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	var ana func(r, c int) bool
	status := make([]int, len(board))
	index := 0
	ana = func(r, c int) bool {
		if index < len(word) {
			var i, j int
			for n := -1; n < 2; n += 2 {
				i = r + n
				j = c
				if i < 0 || i >= len(board) {
					continue
				}

				if status[i]&(1<<j) == 0 &&
					board[i][j] == word[index] {
					index++
					status[i] += 1 << j
					if ana(i, j) {
						return true
					}
					status[i] -= 1 << j
					index--

				}

			}
			for n := -1; n < 2; n += 2 {
				j = c + n
				i = r
				if j < 0 || j >= len(board[0]) {
					continue
				}

				if status[i]&(1<<j) == 0 &&
					board[i][j] == word[index] {
					index++
					status[i] += 1 << j
					if ana(i, j) {
						return true
					}
					status[i] -= 1 << j
					index--

				}

			}
			return false
		} else {
			return true
		}
	}
	for i, row := range board {
		for j, value := range row {
			if value == word[0] {
				index = 1
				status[i] += 1 << j
				if ana(i, j) {
					return true
				}
				status[i] -= 1 << j

			}
		}
	}

	return false
}

// @lc code=end
