/*
 * @lc app=leetcode.cn id=1139 lang=golang
 *
 * [1139] 最大的以 1 为边界的正方形
 */
package Solutions

// @lc code=start
func largest1BorderedSquare(grid [][]int) int {
	var maxGridSize = 0
	var upSize = make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		upSize[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i > 0 {
				upSize[i][j] = upSize[i-1][j] + 1
			} else {
				upSize[i][j] = 1
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			if j > 0 {
				grid[i][j] = grid[i][j-1] + 1
			}
			w := min_len(grid[i][j], upSize[i][j])
			for k := 0; k < w; k++ {
				if grid[i-k][j] >= k+1 && upSize[i][j-k] >= k {
					if maxGridSize < (k+1)*(k+1) {
						maxGridSize = (k + 1) * (k + 1)
					}
				}
			}

		}
	}
	return maxGridSize
}

func min_len(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end
