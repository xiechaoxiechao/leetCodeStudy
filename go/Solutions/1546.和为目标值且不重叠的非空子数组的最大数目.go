/*
 * @lc app=leetcode.cn id=1546 lang=golang
 *
 * [1546] 和为目标值且不重叠的非空子数组的最大数目
 */
package Solutions

// @lc code=start
func maxNonOverlapping(nums []int, target int) int {
	var ans = 0
	var mp = make(map[int]struct{})
	presum := 0
	mp[0] = struct{}{}
	for i := 0; i < len(nums); i++ {
		presum += nums[i]

		if _, ok := mp[presum-target]; ok {
			ans++
			mp = map[int]struct{}{}
		}
		mp[presum] = struct{}{}

	}
	return ans

}

// @lc code=end
