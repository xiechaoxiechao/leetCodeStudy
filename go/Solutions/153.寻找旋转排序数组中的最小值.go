/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
package Solutions

// @lc code=start
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	var ans = nums[0]
	for l <= r {
		mid := l + (r-l)/2
		if nums[l] < nums[mid] {
			if ans > nums[l] {
				ans = nums[l]
			}
			l = mid + 1
		} else if nums[mid] < nums[r] {
			if ans > nums[mid] {
				ans = nums[mid]
			}
			r = mid - 1
		} else {
			if ans > nums[mid] {
				ans = nums[mid]
			}
			l++
		}

	}
	return ans
}

// @lc code=end
