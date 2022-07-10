/*
 * @lc app=leetcode.cn id=795 lang=golang
 *
 * [795] 区间子数组个数
 */

// @lc code=start
package Solutions

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	var l = 0
	var ans = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > right {
			var n = i - l
			ans += (n + 1) * n / 2
			l = i
		}

	}
	return ans

}

// @lc code=end
