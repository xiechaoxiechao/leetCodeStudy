/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package Solutions

// @lc code=start
func rob(nums []int) int {
	var dp = make([]int, 100)
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp[0] = nums[0]
	dp[1] = max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
