/*
 * @lc app=leetcode.cn id=1637 lang=golang
 *
 * [1637] 两点之间不包含任何点的最宽垂直面积
 */
package Solutions

import "sort"

// @lc code=start
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	var ans = 0
	for i := 1; i < len(points); i++ {
		t := points[i][0] - points[i-1][0]
		if ans < t {
			ans = t
		}
	}
	return ans
}

// @lc code=end
