/*
 * @lc app=leetcode.cn id=1433 lang=golang
 *
 * [1433] 检查一个字符串是否可以打破另一个字符串
 */
package Solutions

import (
	"sort"
)

// @lc code=start
func checkIfCanBreak(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Slice(b1, func(i, j int) bool {
		if b1[i] < b1[j] {
			return true
		}
		return false
	})
	sort.Slice(b2, func(i, j int) bool {
		if b2[i] < b2[j] {
			return true
		}
		return false
	})
	flag := false
	i := 0
	for ; i < len(b1); i++ {
		if b1[i] < b2[i] {
			flag = true
			i++
			break
		} else if b1[i] == b2[i] {
			continue
		} else {
			i++
			break
		}

	}
	for ; i < len(s1); i++ {
		if b1[i] < b2[i] != flag && b1[i] != b2[i] {
			return false
		}
	}
	return true
}

// @lc code=end
