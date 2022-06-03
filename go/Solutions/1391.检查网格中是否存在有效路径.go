/*
 * @lc app=leetcode.cn id=1391 lang=golang
 *
 * [1391] 检查网格中是否存在有效路径
 */

// @lc code=start
package Solutions

func hasValidPath(grid [][]int) bool {
	i, j := 0, 0
	var direct = true
	for {
		switch grid[i][j] {
		case 1:
			if direct {
				j += 1
				if j >= len(grid[0]) {
					return false
				}
				if grid[i][j] == 3 {

				} else if grid[i][j] == 5 {
					direct = false
				} else {
					return false
				}
			} else {
				j -= 1
				if j < 0 {
					return false
				}
				if grid[i][j] == 4 {
					direct = true
				} else if grid[i][j] == 6 {

				} else {
					return false
				}
			}
			break

		case 2:
			if direct {
				i++
				if i > len(grid) {
					return false
				}
				if grid[i][j] == 5 {
					direct = false
				} else if grid[i][j] == 6 {

				} else {
					return false
				}
			} else {
				i--
				if i < 0 {
					return false
				}
				if grid[i][j] == 3 {

				} else if grid[i][j] == 4 {
					direct = true
				} else {
					return false
				}
			}

		case 3:
			if direct {
				i++
				if i >= len(grid) {
					return false
				}
				if grid[i][j] == 2 {

				} else if grid[i][j] == 5 {
					direct = false
				} else if grid[i][j] == 6 {

				} else {
					return false
				}
			} else {
				j--
				if j < 0 {
					return false
				}
				if grid[i][j] == 1 {

				} else if grid[i][j] == 4 {
					direct = true
				} else if grid[i][i] == 6 {

				} else {
					return false
				}
			}
		case 4:
			if direct {

			} else {

			}
		case 5:
		case 6:
		}
	}
}

// @lc code=end
