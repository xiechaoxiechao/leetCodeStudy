/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */
package Solutions

// @lc code=start
func longestOnes(nums []int, k int) int {
	var zeroNum = 0
	var n = len(nums)
	var ans = 0
	for l, r := -1, 0; r < n; r++ {
		if nums[r] == 1 {
			if ans < (r - l) {
				ans = r - l
			}
		} else {
			zeroNum++
			if zeroNum > k {
				for l < r {
					l++
					if nums[l] == 0 {
						break
					}
				}
				zeroNum--
			} else {
				if ans < (r - l) {
					ans = r - l
				}
			}
		}
	}
	return ans
}

// @lc code=end
