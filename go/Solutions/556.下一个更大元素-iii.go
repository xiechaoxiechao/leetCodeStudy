/*
 * @lc app=leetcode.cn id=556 lang=golang
 *
 * [556] 下一个更大元素 III
 */

// @lc code=start
package Solutions

import (
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	var s = []byte(strconv.Itoa(n))
	var l = len(s) - 2
	for l >= 0 {
		if s[l] < s[l+1] {
			min := byte(10)
			minI := l
			for i := l + 1; i < len(s); i++ {
				p := s[i] - s[l]
				if p > 0 && p < min {
					min = s[i] - s[l]
					minI = i
				}
			}
			s[l], s[minI] = s[minI], s[l]
			sort.Slice(s[l+1:], func(i, j int) bool {
				return s[l+1+i] < s[l+1+j]
			})
			var res, _ = strconv.Atoi(string(s))
			if res > math.MaxInt32 {
				return -1
			}
			return res

		}
		l--
	}
	return -1
}

// @lc code=end
