/*
 * @lc app=leetcode.cn id=1281 lang=golang
 *
 * [1281] 整数的各位积和之差
 */
package Solutions

// @lc code=start
func subtractProductAndSum(n int) int {
	var mul = 1
	var sum = 0
	for n > 0 {
		t := n % 10
		sum += t
		mul *= t
		n /= 10
	}
	mul -= sum

	return mul
}

// @lc code=end
