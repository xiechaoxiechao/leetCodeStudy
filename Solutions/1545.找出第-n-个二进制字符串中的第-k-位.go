/*
 * @lc app=leetcode.cn id=1545 lang=golang
 *
 * [1545] 找出第 N 个二进制字符串中的第 K 位
 */
package Solutions

// @lc code=start
func findKthBit(n int, k int) byte {
	n = 1 << n
	var flag = false
	for n > 2 {
		if k == (n >> 1) {
			if flag {
				return '0'
			} else {
				return '1'
			}
		}
		if k > (n >> 1) {
			k = n - k
			flag = !flag
		}
		n = n >> 1
	}
	if flag {
		return '1'
	} else {
		return '0'
	}

}

// @lc code=end
