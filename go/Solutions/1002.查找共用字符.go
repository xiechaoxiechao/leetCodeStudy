/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找共用字符
 */

// @lc code=start
package Solutions

func commonChars(words []string) []string {
	var st = make([]int, 26)
	for _, v := range words[0] {
		st[v-'a']++
	}
	for i := 1; i < len(words); i++ {
		var tem = make([]int, 26)
		for _, v := range words[i] {
			tem[v-'a']++
		}
		for j := 0; j < 26; j++ {
			if st[j] > tem[j] {
				st[j] = tem[j]
			}
		}
	}
	var ans = make([]string, 0, 26)

	for i := 0; i < 26; i++ {
		for j := 0; j < st[i]; j++ {
			ans = append(ans, string([]byte{byte(i) + 'a'}))
		}
	}
	return ans
}

// @lc code=end
