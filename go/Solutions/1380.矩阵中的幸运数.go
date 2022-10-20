/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] 矩阵中的幸运数
 */
package Solutions

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	var ans = make([]int, 0, 8)
	for i := 0; i < len(matrix); i++ {
		index := 0
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][index] > matrix[i][j] {
				index = j
			}
		}
		j := 0
		for ; j < len(matrix); j++ {
			if matrix[i][index] < matrix[j][index] {
				break
			}
		}
		if j == len(matrix) {
			ans = append(ans, matrix[i][index])
		}
	}
	return ans
}

// @lc code=end
