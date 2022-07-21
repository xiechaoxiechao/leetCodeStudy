/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package Solutions

// @lc code=start
func numIslands(grid [][]byte) int {
	var gridT = make([][]byte, len(grid)+2)
	var islandCount = 0
	for i := 0; i < len(grid); i++ {
		gridT[i+1] = make([]byte, 1)
		gridT[i+1][0] = '0'
		gridT[i+1] = append(gridT[i+1], grid[i]...)
		gridT[i+1] = append(gridT[i+1], '0')
	}
	gridT[0] = make([]byte, len(grid[0])+2)
	gridT[len(gridT)-1] = make([]byte, len(grid[0])+2)
	for i := 0; i < len(gridT[1]); i++ {
		gridT[0][i] = '0'
		gridT[len(gridT)-1][i] = '0'
	}
	for i := 1; i <= len(grid); i++ {
		for j := 1; j <= len(grid[0]); j++ {
			if gridT[i][j] == '1' {
				islandCount++
				f1(i, j, gridT)
			}
		}
	}
	return islandCount
}

func f1(i, j int, grid [][]byte) {
	grid[i][j] = '0'
	if grid[i-1][j] == '1' {
		f1(i-1, j, grid)
	}
	if grid[i+1][j] == '1' {
		f1(i+1, j, grid)
	}
	if grid[i][j-1] == '1' {
		f1(i, j-1, grid)
	}
	if grid[i][j+1] == '1' {
		f1(i, j+1, grid)
	}

}

// @lc code=end
