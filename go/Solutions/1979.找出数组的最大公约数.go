/*
 * @lc app=leetcode.cn id=1979 lang=golang
 *
 * [1979] 找出数组的最大公约数
 */
package Solutions

// @lc code=start
func findGCD(nums []int) int {
	var maxNum, minNum = -1, 1001
	for i := 0; i < len(nums); i++ {
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
		if minNum > nums[i] {
			minNum = nums[i]
		}
	}
	for i := minNum; i > 0; i-- {
		if maxNum%i == 0 && minNum%i == 0 {
			return i
		}
	}
	return 1
}

// @lc code=end
