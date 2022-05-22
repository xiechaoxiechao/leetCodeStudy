package Solutions

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	var res = make([][]int, 0, 1<<len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var mp = make(map[int]int)
	for _, i := range nums {
		mp[i] += 1
	}
	var be func(int, []int, int)
	be = func(i int, sl []int, index int) {
		var tem = index
		for i < len(nums) {
			var temI = i
			i += mp[nums[i]]
			for j := 0; j < mp[nums[temI]]; j++ {
				sl[index] = nums[temI]
				index++
				var tem = make([]int, index)
				copy(tem, sl)
				res = append(res, tem)
				be(i, sl, index)

			}
			index = tem

		}

	}
	res = append(res, []int{})
	be(0, make([]int, len(nums)), 0)

	return res

}

// @lc code=end
