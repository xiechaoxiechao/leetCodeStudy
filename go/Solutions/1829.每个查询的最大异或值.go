/*
 * @lc app=leetcode.cn id=1829 lang=golang
 *
 * [1829] 每个查询的最大异或值
 */
package Solutions

// @lc code=start
func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	ans := make([]int, n)
	tem := 0
	for i := 0; i < n; i++ {
		tem ^= nums[i]
		k := 0
		for j := 0; j < maximumBit; j++ {
			if tem&(1<<j) == 0 {
				k += (1 << j)
			}
		}
		ans[n-1-i] = k
	}
	return ans
}

// @lc code=end
