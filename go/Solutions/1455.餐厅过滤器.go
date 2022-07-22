/*
 * @lc app=leetcode.cn id=1566 lang=golang
 *
 * [1455] 检查单词是否为句中其他单词的前缀
 */
package Solutions

import "strings"

// @lc code=start
func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	searchWordLen := len(searchWord)
	for ind, word := range words {
		if searchWordLen > len(word) {
			continue
		}
		var i int
		for i = 0; i < searchWordLen; i++ {
			if word[i] != searchWord[i] {
				break
			}
		}
		if i == searchWordLen {
			return ind + 1
		}
	}
	return -1
}

// @lc code=end
