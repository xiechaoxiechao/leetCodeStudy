/*
 * @lc app=leetcode.cn id=868 lang=golang
 *
 * [868] 二进制间距
 */

// @lc code=start
package Solutions

func binaryGap(n int) int {
	var last = 1000000
	var max = 0
	for i := 0; (1 << i) <= n; i++ {
		if n&(1<<i) != 0 {
			if max < i-last {
				max = i - last
			}
			last = i
		}
	}
	return max
}

// @lc code=end
