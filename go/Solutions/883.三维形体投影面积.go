/*
 * @lc app=leetcode.cn id=883 lang=golang
 *
 * [883] 三维形体投影面积
 */

// @lc code=start
package Solutions

func projectionArea(grid [][]int) int {
	var l = len(grid)
	var area = 0
	for i := 0; i < l; i++ {
		max := 0
		max2 := 0
		for j := 0; j < l; j++ {
			if grid[i][j] != 0 {
				area++
			}
			if max < grid[i][j] {
				max = grid[i][j]
			}
			if max2 < grid[j][i] {
				max2 = grid[j][i]
			}
		}
		area += max
		area += max2

	}
	return area
}

// @lc code=end
