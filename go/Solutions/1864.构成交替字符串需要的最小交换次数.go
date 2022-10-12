/*
 * @lc app=leetcode.cn id=1864 lang=golang
 *
 * [1864] 构成交替字符串需要的最小交换次数
 */
package Solutions

// @lc code=start
func minSwaps(s string) int {
	oneCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			oneCount++
		}
	}
	zeroCount := len(s) - oneCount
	if zeroCount-oneCount < -1 || zeroCount-oneCount > 1 {
		return -1
	}
	if zeroCount == oneCount {
		count := 0
		for i := 0; i < len(s); i++ {
			if (s[i] == '1' && i%2 != 0) || (s[i] == '0' && i%2 != 1) {
				count++
			}
		}
		if count > len(s)-count {
			count = len(s) - count
		}
		return count / 2
	}
	if zeroCount > oneCount {
		count := 0
		for i := 0; i < len(s); i++ {
			if (s[i] == '0' && i%2 != 0) || (s[i] == '1' && i%2 != 1) {
				count++
			}
		}
		return count / 2
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if (s[i] == '0' && i%2 != 1) || (s[i] == '1' && i%2 != 0) {
			count++
		}
	}
	return count / 2

}

// @lc code=end
