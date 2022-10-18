/*
 * @lc app=leetcode.cn id=944 lang=golang
 *
 * [944] 删列造序
 */
package Solutions

import "sort"

// @lc code=start
func minDeletionSize(strs []string) int {
	var ans = 0
	for i1 := 0; i1 < len(strs[0]); i1++ {
		if !sort.SliceIsSorted(strs, func(i, j int) bool {
			return strs[i][i1] < strs[j][i1]
		}) {
			ans++
		}
	}
	return ans

}

// @lc code=end
