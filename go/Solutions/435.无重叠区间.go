/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */
package Solutions

import "sort"

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	ans := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	lastEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < lastEnd {
			ans++
		} else {
			lastEnd = intervals[i][1]
		}
	}
	return ans
}

// @lc code=end
