/*
 * @lc app=leetcode.cn id=2164 lang=golang
 *
 * [2164] 对奇偶下标分别排序
 */
package Solutions

import "sort"

// @lc code=start
func sortEvenOdd(nums []int) []int {
	var nums1 []int
	var nums2 []int
	if len(nums)%2 == 0 {
		nums1 = make([]int, len(nums)/2)
		nums2 = make([]int, len(nums)/2)
	} else {
		nums1 = make([]int, len(nums)/2)
		nums2 = make([]int, len(nums)/2+1)
	}
	for i, j := 0, 0; i < len(nums); i += 2 {
		nums2[j] = nums[i]
		j++
	}
	for i, j := 1, 0; i < len(nums); i += 2 {
		nums1[j] = nums[i]
		j++
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 1, len(nums1)-1; i < len(nums); i += 2 {
		nums[i] = nums1[j]
		j--
	}
	for i, j := 0, 0; i < len(nums); i += 2 {
		nums[i] = nums2[j]
		j++
	}
	return nums

}

// @lc code=end
