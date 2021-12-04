package Solutions

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	sTop := 0
	stack := make([]int, len(heights)+1)
	stack[0] = -1
	res := 0
	tem := 0
	for i := 0; i < len(heights); i++ {
		for sTop != 0 && heights[stack[sTop]] > heights[i] {
			tem = heights[stack[sTop]] * (i - stack[sTop-1] - 1)
			if tem > res {
				res = tem
			}
			sTop--
		}
		sTop++
		stack[sTop] = i

	}
	for i := 1; i <= sTop; i++ {
		tem = heights[stack[i]] * (stack[sTop] - stack[i-1])
		if tem > res {
			res = tem
		}

	}
	return res
}

// @lc code=end
