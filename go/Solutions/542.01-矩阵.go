/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */
package Solutions

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	var h, w = len(mat), len(mat[0])
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if mat[i][j] == 0 {
				continue
			}
			mat[i][j] = 1<<31 - 1
			if i-1 >= 0 {
				mat[i][j] = mat[i-1][j] + 1
			}
			if j-1 >= 0 {
				mat[i][j] = min_up(mat[i][j], mat[i][j-1]+1)
			}
		}
	}
	for i := h - 1; i >= 0; i-- {
		for j := w - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}
			if i+1 < h {
				mat[i][j] = min_up(mat[i][j], mat[i+1][j]+1)
			}
			if j+1 < w {
				mat[i][j] = min_up(mat[i][j], mat[i][j+1]+1)
			}
		}
	}
	return mat

}

func min_up(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end
