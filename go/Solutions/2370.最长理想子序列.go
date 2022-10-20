/*
 * @lc app=leetcode.cn id=2370 lang=golang
 *
 * [2370] 最长理想子序列
 */
package Solutions

// @lc code=start
func longestIdealString(s string, k int) int {
	if k == 25 {
		return len(s)
	}
	var lo = [26]int{}
	max := -1
	for i := 0; i < len(s); i++ {
		l := s[i] - byte(k)
		r := s[i] + byte(k)
		if l < 'a' {
			l = 'a'
		}
		if r > 'z' {
			r = 'z'
		}
		l -= 'a'
		r -= 'a'
		max = -1
		for l <= r {
			if max < lo[l] {
				max = lo[l]
			}
			l++
		}
		lo[s[i]-'a'] = max + 1
	}
	var ans = -1
	for i := 0; i < 26; i++ {
		if ans < lo[i] {
			ans = lo[i]
		}
	}
	return ans
}

// @lc code=end
