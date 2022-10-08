/*
 * @lc app=leetcode.cn id=1837 lang=golang
 *
 * [1837] K 进制表示下的各位数字总和
 */
package Solutions

// @lc code=start
func sumBase(n int, k int) int {
	ans := 0
	for n > 0 {

		ans += n % k
		n /= k

	}

	return ans
}

// @lc code=end
