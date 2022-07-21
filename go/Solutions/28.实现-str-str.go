/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */
package Solutions

// @lc code=start
func strStr(haystack string, needle string) int {
	l := len(needle)
	if l == 0 {
		return 0
	}
	next := make([]int, l, l)
	next[0] = -1
	j := 0
	if l == 2 {
		next[1] = 0
	} else {
		for i := 1; i < l-1; {
			if j == -1 || needle[i] == needle[j] {
				j++
				next[i+1] = j
				i++
			} else {
				j = next[j]
			}
		}
	}
	l2 := len(haystack)
	j = 0
	res := -1
	for i := 0; i < l2; {
		if j == -1 || haystack[i] == needle[j] {
			if j == l-1 {
				res = i - l + 1
				break
			}
			j++
			i++
		} else {
			j = next[j]
		}
	}
	return res
}

// @lc code=end
