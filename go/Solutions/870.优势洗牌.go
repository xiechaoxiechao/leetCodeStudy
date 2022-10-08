/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 */
package Solutions

import "sort"

type item struct {
	ind int
	val int
}

// @lc code=start
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	sort.Ints(nums1)
	var cpNm2 = make([]item, n)
	for i := 0; i < n; i++ {
		cpNm2[i].ind = i
		cpNm2[i].val = nums2[i]
	}

	sort.Slice(cpNm2, func(i, j int) bool {
		return cpNm2[i].val < cpNm2[j].val
	})
	var ans = make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	for i, j := 0, 0; i < n && j < n; {
		if nums1[i] <= cpNm2[j].val {
			i++
		} else {
			ans[cpNm2[j].ind] = nums1[i]
			nums1[i] = -1
			j++
			i++
		}
	}
	for i, j := 0, 0; i < n; {
		if nums1[i] == -1 {
			i++
			continue
		}
		if ans[j] != -1 {
			i++
			continue
		}
		ans[j] = ans[i]
		i++
		j++
	}
	return ans
}

// @lc code=end
