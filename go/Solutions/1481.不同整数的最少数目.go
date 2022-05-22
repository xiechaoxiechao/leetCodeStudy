/*
 * @lc app=leetcode.cn id=1481 lang=golang
 *
 * [1481] 不同整数的最少数目
 */
package Solutions

import "sort"

// @lc code=start

func findLeastNumOfUniqueInts(arr []int, k int) int {
	ma := make(map[int]int)
	for _, v := range arr {
		ma[v] = ma[v] + 1
	}
	sl := make([]int, len(ma))
	i := 0
	for _, value := range ma {
		sl[i] = value
		i++
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	i = 0
	for k > 0 {
		if sl[i] <= k {
			k -= sl[i]
			i++
		} else {
			break
		}
	}
	return len(sl) - i

}

// @lc code=end
