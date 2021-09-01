package Solutions

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		if nums[0] == target {
			res[0] = 0
			res[1] = 0
		}
		return res
	}
	l, r, mid := 0, len(nums)-1, 0
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		res[0] = l
	} else {
		return res
	}
	r = len(nums) - 1
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	res[1] = l
	return res

}

// @lc code=end
