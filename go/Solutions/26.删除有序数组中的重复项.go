package Solutions

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	for i := 1; i < len(nums); {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
			n--
			continue
		}
		i++
	}
	return n

}

// @lc code=end
