/*
 * @lc app=leetcode.cn id=1329 lang=golang
 *
 * [1329] 将矩阵按对角线排序
 */
package Solutions

import "sort"

// @lc code=start
func diagonalSort(mat [][]int) [][]int {

	m := len(mat)
	n := len(mat[0])
	var temSlice = make([]int, _max(m, n))
	for i := 0; i < m; i++ {
		maxL := maxLen(i, 0, m, n)
		for j := 0; j < maxL; j++ {
			temSlice[j] = mat[i+j][j]
		}
		sort.Ints(temSlice[0:maxL])
		for j := 0; j < maxL; j++ {
			mat[i+j][j] = temSlice[j]
		}
	}

	for i := 1; i < n; i++ {
		maxL := maxLen(0, i, m, n)
		for j := 0; j < maxL; j++ {
			temSlice[j] = mat[j][i+j]
		}
		sort.Ints(temSlice[0:maxL])
		for j := 0; j < maxL; j++ {
			mat[j][j+i] = temSlice[j]
		}
	}
	return mat
}
func _max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxLen(a, b, m, n int) int {
	l1 := m - a
	l2 := n - b
	if l1 < l2 {
		return l1
	}
	return l2
}

// @lc code=end
