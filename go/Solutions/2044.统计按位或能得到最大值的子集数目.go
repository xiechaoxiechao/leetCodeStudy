/*
 * @lc app=leetcode.cn id=2044 lang=golang
 *
 * [2044] 统计按位或能得到最大值的子集数目
 */

// @lc code=start
package Solutions

func countMaxOrSubsets(nums []int) int {
	var dfs func(ind, or int)
	var maxOr = -1
	var maxNum = 0
	dfs = func(ind, or int) {
		for i := ind; i < len(nums); i++ {
			tem := or | nums[i]
			if tem > maxOr {
				maxOr = tem
				maxNum = 1
			} else if tem == maxOr {
				maxNum++
			}
			dfs(i+1, tem)
		}
	}
	dfs(0, 0)
	return maxNum
}

// @lc code=end
