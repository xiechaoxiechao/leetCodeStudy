package Solutions

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	// nums[left] = left + 1
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == left+1 {
			left += 1
		} else if nums[left] < left+1 || nums[left] > right || nums[left] == nums[nums[left]-1] {
			// 出现缺失
			nums[left], nums[right-1] = nums[right-1], nums[left]
			right--
		} else {
			index := nums[left]
			nums[left], nums[index-1] = nums[index-1], nums[left]
		}
	}
	return left + 1
}

// @lc code=end
