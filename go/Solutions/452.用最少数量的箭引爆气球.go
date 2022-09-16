/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */
package Solutions

import "sort"

// @lc code=start
func findMinArrowShots(points [][]int) int {
	ans := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	lastEnd := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > lastEnd {
			ans++
			lastEnd = points[i][1]
		}
	}
	return ans
}

// @lc code=end
