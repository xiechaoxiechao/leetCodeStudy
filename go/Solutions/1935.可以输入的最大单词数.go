/*
 * @lc app=leetcode.cn id=1935 lang=golang
 *
 * [1935] 可以输入的最大单词数
 */
package Solutions

// @lc code=start
func canBeTypedWords(text string, brokenLetters string) (ans int) {
	mp := make(map[byte]struct{}, len(brokenLetters))
	for i := 0; i < len(brokenLetters); i++ {
		mp[brokenLetters[i]] = struct{}{}
	}
	text += " "
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			ans++
			continue
		}
		if _, ok := mp[text[i]]; ok {
			for i < len(text) && text[i] != ' ' {
				i++
			}
		}
	}
	return

}

// @lc code=end
