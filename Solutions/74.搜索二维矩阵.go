package Solutions

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix) - 1
	if right == -1 || target < matrix[0][0] || target > matrix[right][len(matrix[0])-1] {
		return false
	}
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if target <= matrix[mid][len(matrix[0])-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	tem := left
	left = 0
	right = len(matrix[0]) - 1
	for left < right {
		mid = left + (right-left)/2
		if target <= matrix[tem][mid] {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return matrix[tem][left] == target
}

// @lc code=end
