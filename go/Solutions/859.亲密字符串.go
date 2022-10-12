/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */
package Solutions

// @lc code=start
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	i := -1
	j := -1
	for k := 0; k < len(s); k++ {
		if s[k] != goal[k] {
			if i == -1 {
				i = k
				continue
			} else if j == -1 {
				j = k
				continue
			} else {
				return false
			}

		}
	}
	if i == -1 && j == -1 {
		mp := make(map[byte]struct{})
		for i := 0; i < len(s); i++ {
			if _, ok := mp[s[i]]; ok {
				return true
			} else {
				mp[s[i]] = struct{}{}
			}
		}
	}
	if i >= 0 && j >= 0 && s[i] == goal[j] && s[j] == goal[i] {
		return true
	} else {
		return false
	}
}

// @lc code=end
