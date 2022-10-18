/*
 * @lc app=leetcode.cn id=2351 lang=golang
 *
 * [2351] 第一个出现两次的字母
 */
package Solutions

// @lc code=start
func repeatedCharacter(s string) byte {
	var mp = [26]bool{}
	for i := 0; i < len(s); i++ {
		if mp[s[i]-'a'] {
			return s[i]
		} else {
			mp[s[i]-'a'] = true
		}
	}
	return 'a'

}

// @lc code=end
