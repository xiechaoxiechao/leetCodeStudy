/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
package Solutions

import (
	"sort"
)

type node struct {
	val int
	ind int
}

func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var indexs = make([]node, len(nums))
	for i := 0; i < len(nums); i++ {
		indexs[i].val = nums[i]
		indexs[i].ind = i
	}

	sort.Slice(indexs[1:len(indexs)-1], func(i, j int) bool {
		return indexs[i].val > indexs[j].val
	})
	var ans = 0
	var status = make([]int, len(nums))
	for i := 1; i < len(indexs)-1; i++ {
		var l = indexs[i].ind
		var tem = nums[l]
		status[l] = 1

		for status[l] == 1 {
			l--
		}
		tem *= nums[l]
		l = indexs[i].ind
		for status[l] == 1 {
			l++
		}
		tem *= nums[l]
		ans += tem

	}
	ans += nums[0] * nums[len(nums)-1]
	if nums[0] < nums[len(nums)-1] {
		ans += nums[len(nums)-1]
	} else {
		ans += nums[0]
	}
	return ans

}

// @lc code=end
