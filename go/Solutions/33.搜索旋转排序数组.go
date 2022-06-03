package Solutions

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		if nums[0] != target {
			return -1
		} else {
			return 0
		}
	}
	l, r, mid := 0, len(nums)-1, 0
	for {
		if r-l <= 2 {
			for i := l; i <= r; i++ {
				if nums[i] == target {
					return i
				}
			}
			return -1
		}
		if nums[l] < nums[r] {
			mid = l + (r-l)/2
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				r = mid
			} else {
				l = mid
			}
		} else {
			mid = l + (r-l)/2
			if nums[l] < nums[mid] {
				if nums[l] <= target && nums[mid] >= target {
					r = mid
				} else {
					l = mid
				}
			} else {
				if nums[mid] <= target && nums[r] >= target {
					l = mid
				} else {
					r = mid
				}
			}

		}
	}

}

// @lc code=end
