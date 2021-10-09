package Solutions

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	if len(heights) == 1 {
		return heights[0]
	}
	s1 := make([]int, len(heights))
	s2 := make([]int, len(heights))
	s1[0] = 0
	res := 0
	j := 0
	for i := 1; i < len(heights); i++ {
		if heights[i] < heights[i-1] {
			for j = s1[i-1] - 1; j >= 0 && heights[j] >= heights[i]; j-- {
			}
			s1[i] = j + 1
		} else if heights[i] == heights[i-1] {
			s1[i] = s1[i-1]
		} else {
			s1[i] = i
		}
	}
	s2[len(heights)-1] = len(heights) - 1
	for i := len(heights) - 2; i > -1; i-- {
		if heights[i] < heights[i+1] {
			for j = s2[i+1] + 1; j < len(heights) && heights[j] >= heights[i]; j++ {
			}
			s2[i] = j - 1
		} else if heights[i] == heights[i+1] {
			s2[i] = s2[i+1]
		} else {
			s2[i] = i
		}
	}
	var tem = 0
	for i, v := range s1 {
		tem = (s2[i] - v + 1) * heights[i]
		if res < tem {
			res = tem
		}
	}
	return res
}

// @lc code=end
