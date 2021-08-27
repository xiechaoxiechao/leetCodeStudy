package Solutions

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	min := []int{9, 9, 9, 9, 9, 9, 9, 9, 9}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				min[i]--
			}
		}
	}
	fill(board, min)
}

func fill(b [][]byte, min []int) bool {
	index := 0
	tem := 10
	for i, va := range min {
		if va != 0 && va < tem {
			tem = va
			index = i
		}
	}
	min[index]--
	if tem == 10 {
		return true
	}
	cl := 0
	for i := 0; i < 9; i++ {
		if b[index][i] == '.' {
			cl = i
			break
		}
	}
	var i byte
	ofC := (cl / 3) * 3
	ofR := (index / 3) * 3

	for i = 49; i < 58; i++ {
		j := 0
		for ; j < 9; j++ {
			if b[index][j] == i || b[j][cl] == i {
				break
			}
			if j < 3 {
				if b[ofR][ofC+j] == i {
					break
				}

			} else if j < 6 {
				if b[ofR+1][ofC+j-3] == i {
					break
				}
			} else {
				if b[ofR+2][ofC+j-6] == i {
					break
				}
			}
		}
		if j != 9 {
			continue
		}
		b[index][cl] = i
		if fill(b, min) {
			return true
		} else {
			b[index][cl] = '.'
		}
	}
	min[index]++
	return false
}

// @lc code=end
