/*
 * @lc app=leetcode.cn id=1957 lang=golang
 *
 * [1957] 删除字符使字符串变好
 */
package Solutions

// @lc code=start
func makeFancyString(s string) string {
	lastNum := 1
	ans := make([]byte, 0, len(s))
	ans = append(ans, s[0])
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			if lastNum == 2 {
				continue
			} else {
				lastNum++
				ans = append(ans, s[i])
			}
		} else {
			lastNum = 1
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

// @lc code=end
