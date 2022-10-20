/*
 * @lc app=leetcode.cn id=2231 lang=golang
 *
 * [2231] 按奇偶性交换后的最大数字
 */
package Solutions

import "sort"

// @lc code=start
func largestInteger(num int) int {
	var bit = make([]bool, 0, 10)
	var a = make([]int, 0, 10)
	var b = make([]int, 0, 10)

	for num > 0 {
		t := num % 10
		if t%2 == 0 {
			bit = append(bit, true)
			a = append(a, t)
		} else {
			bit = append(bit, false)
			b = append(b, t)
		}
		num /= 10
	}
	sort.Ints(a)
	sort.Ints(b)
	var p1 = 0
	var p2 = 0
	mul := 1
	var ans = 0
	for i := 0; i < len(bit); i++ {
		if bit[i] {
			ans += a[p1] * mul
			p1++
		} else {
			ans += b[p2] * mul
			p2++
		}
		mul *= 10
	}
	return ans

}

// @lc code=end
