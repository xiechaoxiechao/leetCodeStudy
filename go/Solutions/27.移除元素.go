package Solutions

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	l := len(nums)
	var ex int
	for i := 0; i < l; i++ {
		if nums[i] == val {
			ex = nums[i]
			nums[i] = nums[l-1]
			nums[l-1] = ex
			l--
			i--
			continue
		}
	}
	return l
}

// @lc code=end
