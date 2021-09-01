package Solutions

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	l := len(words)
	res := make([]int, 0)
	if l == 0 {
		return res
	}
	var ok bool
	aMap := make(map[string]int)
	for i := 0; i < l; i++ {
		_, ok := aMap[words[i]]
		if ok {
			aMap[words[i]]++
		} else {
			aMap[words[i]] = 1
		}

	}
	wL := len(words[0])
	index := 0
	sMap := make(map[string]int)
	for ; index <= len(s)-l*wL; index++ {
		for i := 0; i < l; i++ {
			sMap[words[i]] = 0
		}
		for i := 0; i < l*wL; i += wL {
			_, ok = sMap[s[index+i:index+wL+i]]
			if !ok {
				break
			} else {
				sMap[s[index+i:index+wL+i]]++
				if i == (l-1)*wL {
					sw := 0
					for key, value := range sMap {
						if value != aMap[key] {
							sw = 1
						}
					}
					if sw == 0 {
						res = append(res, index)
					}
				}

			}
		}

	}
	return res
}

// @lc code=end
