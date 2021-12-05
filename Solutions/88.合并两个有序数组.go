package Solutions

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	var tem = make([]int, m, m)
	copy(tem, nums1)
	l1 := 0
	l2 := 0
	head := 0
	for l1 < m && l2 < n {
		if tem[l1] < nums2[l2] {
			nums1[head] = tem[l1]
			head++
			l1++
		} else {
			nums1[head] = nums2[l2]
			head++
			l2++
		}
	}
	for ; l1 < m; l1++ {
		nums1[head] = tem[l1]
		head++

	}
	for ; l2 < n; l2++ {
		nums1[head] = nums2[l2]
		head++
	}
}

// @lc code=end
