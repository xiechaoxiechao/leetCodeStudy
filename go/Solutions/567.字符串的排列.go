/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */
package Solutions

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	var mp = [26]int{}
	for i := 0; i < len(s1); i++ {
		mp[s1[i]-'a']++
	}
	var check = func() bool {
		or := 0
		for i := 0; i < 26; i++ {
			or |= mp[i]
		}
		if or == 0 {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < len(s1); i++ {
		mp[s2[i]-'a']--
	}
	if check() {
		return true
	}
	var ls1 = len(s1)
	for i := ls1; i < len(s2); i++ {
		mp[s2[i-ls1]-'a']++
		mp[s2[i]-'a']--
		if check() {
			return true
		}
	}
	return false

}

// @lc code=end
