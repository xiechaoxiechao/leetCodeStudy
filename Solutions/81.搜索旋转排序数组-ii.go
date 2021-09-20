package Solutions

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search1(nums []int, target int) bool {
	mid := 0
	var searchInSorted func(left, right int) bool
	var searchInUnsorted func(left, right int) bool
	searchInSorted = func(left, right int) bool {
		for left < right-1 {
			mid = (left + right) / 2
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid
			}
		}
		if nums[left] == target || nums[right] == target {
			return true
		} else {
			return false
		}

	}
	searchInUnsorted = func(left, right int) bool {
		if right-left <= 1 {
			if nums[left] == target || nums[right] == target {
				return true
			} else {
				return false
			}
		}
		mid = (left + right) / 2
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			return searchInUnsorted(left, mid) || searchInUnsorted(mid, right)
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				return searchInSorted(left, mid)
			} else {
				return searchInUnsorted(mid, right)
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				return searchInSorted(mid, right)
			} else {
				return searchInUnsorted(left, mid)
			}
		}
	}
	return searchInUnsorted(0, len(nums)-1)
}

// @lc code=end
