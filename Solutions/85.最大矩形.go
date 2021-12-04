package Solutions

/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}

		}
		tem := f(heights)
		if tem > max {
			max = tem
		}
	}

	return max
}
func f(heights []int) int {
	sTop := 0
	stack := make([]int, len(heights)+1)
	stack[0] = -1
	res := 0
	for i := 0; i < len(heights); i++ {
		for sTop != 0 && heights[stack[sTop]] > heights[i] {
			tem := heights[stack[sTop]] * (i - stack[sTop-1] - 1)
			if tem > res {
				res = tem
			}
			sTop--
		}

		sTop++
		stack[sTop] = i

	}
	for i := 1; i <= sTop; i++ {
		tem := heights[stack[i]] * (stack[sTop] - stack[i-1])
		if tem > res {
			res = tem
		}

	}
	return res
}

// @lc code=end
