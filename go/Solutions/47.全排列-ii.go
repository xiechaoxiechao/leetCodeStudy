/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */
package Solutions

// @lc code=start
func permuteUnique(nums []int) [][]int {
	length := len(nums)
	var f func(depth int)
	var ans = make([][]int, 0, 16)
	f = func(depth int) {
		if depth == length {
			t := make([]int, length)
			copy(t, nums)
			ans = append(ans, t)
		}
		mp := make(map[int]bool)
		for i := depth; i < length; i++ {
			if !mp[nums[i]] {
				mp[nums[i]] = true
				nums[i], nums[depth] = nums[depth], nums[i]
				f(depth + 1)
				nums[i], nums[depth] = nums[depth], nums[i]
			}
		}
	}
	f(0)
	return ans
}

// @lc code=end
