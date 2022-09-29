/*
 * @lc app=leetcode.cn id=2273 lang=golang
 *
 * [2273] 移除字母异位词后的结果数组
 */
package Solutions

// @lc code=start
func removeAnagrams(words []string) []string {
	var mp1 = &[26]int{0}
	var mp2 = &[26]int{0}
	var ans = make([]string, 0, 8)
	ans = append(ans, words[0])
	for i := 0; i < len(words[0]); i++ {
		mp1[words[0][i]-'a']++
	}
	for i := 1; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			mp2[words[i][j]-'a']++
		}
		if *mp1 != *mp2 {
			ans = append(ans, words[i])
			mp1, mp2 = mp2, mp1
		}
		for k := 0; k < 26; k++ {
			mp2[k] = 0
		}
	}
	return ans

}

// @lc code=end
