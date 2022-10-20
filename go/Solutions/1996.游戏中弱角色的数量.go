/*
 * @lc app=leetcode.cn id=1996 lang=golang
 *
 * [1996] 游戏中弱角色的数量
 */
package Solutions

import "sort"

// @lc code=start
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] < properties[j][0] {
			return true
		} else if properties[i][0] == properties[j][0] {
			return properties[j][1] < properties[i][1]
		} else {
			return false
		}
	})
	var stack = make([][]int, len(properties))
	var st = -1
	var ans = 0
	for i := 0; i < len(properties); {
		if st >= 0 && stack[st][1] < properties[i][1] {
			st--
			ans++
		} else {
			st++
			stack[st] = properties[i]
			i++
		}

	}
	return ans

}

// @lc code=end
