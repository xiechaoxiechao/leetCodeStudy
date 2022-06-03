/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
package Solutions

func firstUniqChar(s string) int {
	var mp = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]] = mp[s[i]] + 1
	}
	for i := 0; i < len(s); i++ {
		if mp[s[i]] == 1 {
			return i
		}
	}

	return -1
}

// @lc code=end
