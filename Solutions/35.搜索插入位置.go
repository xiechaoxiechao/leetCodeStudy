package Solutions

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	l, r, mid := 0, len(nums), 0
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// @lc code=end
