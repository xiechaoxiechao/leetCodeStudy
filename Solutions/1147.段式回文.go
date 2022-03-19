/*
 * @lc app=leetcode.cn id=1147 lang=golang
 *
 * [1147] 段式回文
 */

// @lc code=start
package Solutions

func longestDecomposition(text string) int {
	if len(text) == 1 {
		return 1
	}
	var ans = 1
	var length = 0
	for i, j := 0, len(text)-1; i < j; j-- {
		if text[i] != text[j] {
			length++
		} else {
			flag := true
			for t, k := j+1, 1; k <= length; k, t = k+1, t+1 {
				if text[i+k] != text[t] {
					flag = false
				}
			}
			if flag {
				ans += 2
				i += length + 1
				length = 0
				if i-1 == j-1 {
					ans--
				}
			} else {
				length++
			}
		}
	}
	return ans
}

// @lc code=end
