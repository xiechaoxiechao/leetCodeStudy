/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
package Solutions

func findRepeatedDnaSequences(s string) []string {

	var l = len(s) - 10
	var mp = make(map[string]int)
	for i := 0; i < l; i++ {

		st := s[i : i+10]
		mp[st] += 1
	}
	var ans = make([]string, 0, 100)
	for k, v := range mp {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

// @lc code=end
