/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
package Solutions

import "sort"

func longestWord(words []string) string {
	sort.Strings(words)
	var ans string
	var maxLen = -1
	var set = make(map[string]bool)
	set[""] = true
	for i := 0; i < len(words); i++ {
		if b, ok := set[words[i][:len(words[i])-1]]; ok {
			if len(words[i]) > maxLen && b {
				maxLen = len(words[i])
				ans = words[i]

			}
			set[words[i]] = b
		} else {
			set[words[i]] = false
		}

	}
	return ans
}

// @lc code=end
