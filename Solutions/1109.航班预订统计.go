/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 */

// @lc code=start
package Solutions

func corpFlightBookings(bookings [][]int, n int) []int {
	var c = make([]int, n+1)
	for i := 0; i < len(bookings); i++ {
		c[bookings[i][0]-1] += bookings[i][2]
		c[bookings[i][1]] -= bookings[i][2]
	}
	for i := 1; i < n; i++ {
		c[i] += c[i-1]
	}
	return c[:n]
}

// @lc code=end
