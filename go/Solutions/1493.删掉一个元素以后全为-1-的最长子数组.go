/*
 * @lc app=leetcode.cn id=1493 lang=golang
 *
 * [1493] 删掉一个元素以后全为 1 的最长子数组
 */
package Solutions

// @lc code=start
func longestSubarray(nums []int) int {
	maxLen := 0
	lastLen := 0
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			t := l + lastLen
			if t > maxLen {
				maxLen = t
			}
			lastLen = l
			l = 0

		} else {
			l++
		}
	}
	if l+lastLen > maxLen {
		maxLen = l + lastLen
	}
	if l == len(nums) {
		return l - 1
	}
	return maxLen
}

// @lc code=end
