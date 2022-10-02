/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */
package Solutions

// @lc code=start
func partitionLabels(s string) []int {
	mp := [26]int{0}
	for i := 0; i < len(s); i++ {

		mp[s[i]-'a'] = i
	}
	l := 0
	var ans = make([]int, 0, 8)
	for i := mp[s[l]-'a']; i < len(s); {
		var j int
		for j = l; j < i; j++ {
			i = max_part(i, mp[s[j]-'a'])
		}

		ans = append(ans, j-l+1)

		j++
		l = j

		if l >= len(s) {
			break
		} else {
			i = mp[s[l]-'a']
		}
	}

	return ans
}
func max_part(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end
