/*
 * @lc app=leetcode.cn id=553 lang=golang
 *
 * [553] 最优除法
 */
package Solutions

import (
	"strconv"
	"strings"
)

// @lc code=start

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}
	var res = strings.Builder{}
	res.WriteString(strconv.Itoa(nums[0]))
	res.WriteRune('/')
	res.WriteRune('(')
	res.WriteString(strconv.Itoa(nums[1]))
	for i := 2; i < len(nums); i++ {
		res.WriteRune('/')
		res.WriteString(strconv.Itoa(nums[i]))
	}
	res.WriteRune(')')
	return res.String()

}

// @lc code=end
