/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */
package Solutions

// @lc code=start
func nextGreaterElement_I(nums1 []int, nums2 []int) []int {
	var mp = make(map[int]int, len(nums2))
	for i, v := range nums2 {
		mp[v] = i
	}
	var ans = make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		j := mp[nums1[i]]
		for ; j < len(nums2); j++ {
			if nums1[i] < nums2[j] {
				break
			}
		}
		if j == len(nums2) {
			ans[i] = -1
		} else {
			ans[i] = nums2[j]
		}

	}
	return ans
}

// @lc code=end
