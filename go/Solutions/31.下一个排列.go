package Solutions

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	l := len(nums)
	i := l - 1
	for ; i-1 >= 0; i-- {
		if nums[i] > nums[i-1] {
			j := i
			for ; j < l; j++ {
				if nums[j] <= nums[i-1] {
					break
				}
			}
			nums[i-1], nums[j-1] = nums[j-1], nums[i-1]

			n := i
			m := l - 1
			for m > n {
				nums[n], nums[m] = nums[m], nums[n]
				n++
				m--
			}
			return
		}
	}
	n := 0
	m := l - 1
	for m > n {
		nums[n], nums[m] = nums[m], nums[n]
		n++
		m--
	}
	return

}

// @lc code=end
