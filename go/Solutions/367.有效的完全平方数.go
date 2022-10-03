/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */
package Solutions

// @lc code=start
func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square < num {
			left = mid + 1
		} else if square > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

// @lc code=end
