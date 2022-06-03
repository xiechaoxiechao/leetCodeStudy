/*
 * @lc app=leetcode.cn id=954 lang=golang
 *
 * [954] 二倍数对数组
 */

// @lc code=start
package Solutions

import "sort"

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	if len(arr) == 0 {
		return true
	}
	zero := sort.SearchInts(arr, 0)
	if zero%2 == 1 {
		return false
	}
	var status = make([]int, len(arr))
	for l, r := zero, zero+1; l < len(arr); l++ {
		if status[l] == 1 {
			continue
		}

		tw := arr[l] * 2
		for ; r < len(arr) && arr[r] == tw; r++ {
		}
		if r == len(arr) {
			return false
		}
		status[r] = 1
		r++
	}
	for l, r := zero-1, zero-2; l >= 0; l-- {
		if status[l] == 1 {
			continue
		}

		tw := arr[l] * 2
		for ; r >= 0 && arr[r] == tw; r-- {
		}
		if r == -1 {
			return false
		}
		status[r] = 1
		r--
	}
	return true

}

// @lc code=end
