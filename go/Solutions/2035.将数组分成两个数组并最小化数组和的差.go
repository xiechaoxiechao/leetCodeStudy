/*
 * @lc app=leetcode.cn id=2035 lang=golang
 *
 * [2035] 将数组分成两个数组并最小化数组和的差
 */

// @lc code=start
package Solutions

import (
	"sort"
)

func minimumDifference(nums []int) int {
	var n = len(nums) / 2
	var X = make([][]int, (n+1)/2+1)
	var Y = make([][]int, (n+1)/2+1)
	for mask := 0; mask <= 1<<n; mask++ {
		count := 0
		var x = 0
		var y = 0
		for j := 0; j < n; j++ {
			if (mask>>j)&1 == 1 {
				count++
				x += nums[j]
				y += nums[j+n]
			} else {
				x -= nums[j]
				y -= nums[j+n]
			}

		}
		if count < len(X) {
			X[count] = append(X[count], x)
			Y[count] = append(Y[count], y)
		}

	}
	for i := 0; i < len(X); i++ {
		sort.Ints(Y[i])
	}
	var minSum = 1 << 31

	for i := 0; i < len(X); i++ {
		for j := 0; j < len(X[i]); j++ {
			var y = sort.SearchInts(Y[i], X[i][j])
			if y > 0 {
				minSum = min(minSum, abs(X[i][j], Y[i][y-1]))
			}
			if y < len(Y[i]) {
				minSum = min(minSum, abs(X[i][j], Y[i][y]))
			}

		}
	}
	return minSum

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// @lc code=end
