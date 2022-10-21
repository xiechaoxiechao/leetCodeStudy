/*
 * @lc app=leetcode.cn id=1498 lang=golang
 *
 * [1498] 满足条件的子序列数目
 */
package Solutions

import (
	"sort"
)

// @lc code=start
func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	var ans = 0
	var f = make([]int, len(nums))
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		f[i] = f[i-1] * 2 % (1e9 + 7)
	}
	for i := 0; i < len(nums); {
		if nums[i]*2 > target {
			break
		}
		j := sort.SearchInts(nums, target-nums[i])
		for j < len(nums) && nums[i]+nums[j] == target {
			j++
		}
		ans = (ans + f[j-i-1]) % (1e9 + 7)
		i++

	}
	return ans
}

// @lc code=end
