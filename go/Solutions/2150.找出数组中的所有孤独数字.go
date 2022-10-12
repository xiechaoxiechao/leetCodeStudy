/*
 * @lc app=leetcode.cn id=2150 lang=golang
 *
 * [2150] 找出数组中的所有孤独数字
 */
package Solutions

// @lc code=start
func findLonely(nums []int) []int {
	var mp = make(map[int]int, len(nums))
	for _, v := range nums {
		mp[v]++
	}
	var ans = make([]int, 0, 32)
	for _, v := range nums {
		if mp[v] == 1 && mp[v-1] == 0 && mp[v+1] == 0 {
			ans = append(ans, v)
		}
	}
	return ans
}

// @lc code=end
