package Solutions

import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var left, mid, right, diff int
	res := make([][]int, 0, 1)
	for i := 0; i < length-3; {

		for left = i + 1; left < length-2; {
			mid = left + 1
			right = length - 1
			for mid != right {
				diff = target - (nums[i] + nums[left] + nums[mid] + nums[right])
				if diff == 0 {
					res = append(res, []int{nums[i], nums[left], nums[mid], nums[right]})
					if nums[right] != nums[right-1] {
						right--
					} else {
						for ; nums[right] == nums[right-1] && right != mid; right-- {

						}
						if right == mid {
							break
						} else {
							right--
						}
					}

					continue
				} else if diff < 0 {
					right--
				} else {
					mid++
				}
			}

			if nums[left+1] != nums[left] {
				left++
			} else {
				for ; nums[left+1] == nums[left] && (left+1) < length-1; left++ {
				}
				left++
			}
		}
		if nums[i+1] != nums[i] {
			i++
		} else {
			for ; nums[i+1] == nums[i] && (i+1) < length-1; i++ {
			}
			i++
		}

	}
	// fmt.Println(res)
	return res

}

// @lc code=end
