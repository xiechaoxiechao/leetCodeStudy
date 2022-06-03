/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
package Solutions

func trailingZeroes(n int) (ans int) {
	for n > 0 {
		n /= 5
		ans += n
	}
	return
}

// @lc code=end
