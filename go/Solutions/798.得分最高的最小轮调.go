/*
 * @lc app=leetcode.cn id=798 lang=golang
 *
 * [798] 得分最高的最小轮调
 */

// @lc code=start
package Solutions

func bestRotation(nums []int) int {
	var n = len(nums)
	var ke = make([]int, n)
	for i := 0; i < n; i++ {
		ke[(i-nums[i]+1+n)%n]++
	}
	var res = 0
	for i := 1; i < n; i++ {
		ke[i] += ke[i-1] - 1
		if ke[res] > ke[i] {
			res = i
		}
	}
	return res
}

// @lc code=end
