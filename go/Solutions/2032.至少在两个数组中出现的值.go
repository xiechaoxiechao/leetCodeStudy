/*
 * @lc app=leetcode.cn id=2032 lang=golang
 *
 * [2032] 至少在两个数组中出现的值
 */
package Solutions

// @lc code=start
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var mp = map[int]struct{}{}
	var ans = map[int]struct{}{}
	for _, v := range nums1 {
		mp[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := mp[v]; ok {
			ans[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		mp[v] = struct{}{}
	}
	for _, v := range nums3 {
		if _, ok := mp[v]; ok {
			ans[v] = struct{}{}
		}
	}
	ansSlice := make([]int, 0, len(ans))
	for k, _ := range ans {
		ansSlice = append(ansSlice, k)
	}
	return ansSlice
}

// @lc code=end
