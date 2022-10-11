/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */
package Solutions

// @lc code=start
func setZeroes(matrix [][]int) {
	n := len(matrix) - 1
	m := len(matrix[0]) - 1
	f1 := false
	f2 := false
	for i := 0; i <= n; i++ {
		if matrix[i][0] == 0 {
			f1 = true
			break
		}
	}
	for i := 0; i <= m; i++ {
		if matrix[0][i] == 0 {
			f2 = true
			break
		}
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n; i > 0; i-- {
		for j := m; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if f1 {
		for i := 0; i <= n; i++ {
			matrix[i][0] = 0
		}
	}
	if f2 {
		for i := 0; i <= m; i++ {
			matrix[0][i] = 0
		}
	}

}

// 0  0  3 0
// 0  0  7 8
// 0 10 11 12
// 0 14 15 0

// @lc code=end
