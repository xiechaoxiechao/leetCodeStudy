/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */
package Solutions

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var aim = -1
	var status = make([]bool, len(wordList))
	var res = make([][]string, 0)
	var t = make([]string, len(wordList)+1)
	for i, v := range wordList {
		if v == endWord {
			aim = i
			break
		}
	}

	if aim == -1 {
		return res
	}
	var turnable = make([][]int, len(wordList))
	for i := 0; i < len(wordList); i++ {
		turnable[i] = make([]int, len(wordList)+1)
	}
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			var tem = true
			k := 0
			for ; k < len(beginWord); k++ {
				if wordList[i][k] != wordList[j][k] {
					if tem {
						tem = false
					} else {
						break
					}
				}

			}
			if k == len(beginWord) {
				turnable[i][0]++
				turnable[i][turnable[i][0]] = j
				turnable[j][0]++
				turnable[j][turnable[j][0]] = i
			}

		}
	}
	var l = 0
	var minLen = 20
	var f func(i int)
	f = func(i int) {
		for j := 1; j <= turnable[i][0]; j++ {
			if turnable[i][j] == aim {
				tem := make([]string, l+1)
				copy(tem, t[0:l])
				tem[len(tem)-1] = endWord
				if minLen > len(tem) {
					minLen = len(tem)
					res = res[0:0]
				}
				res = append(res, tem)
				return
			}
		}
		for j := 1; j <= turnable[i][0]; j++ {
			if !status[turnable[i][j]] {
				status[turnable[i][j]] = true
				t[l] = wordList[turnable[i][j]]
				l++
				if l < minLen {
					f(turnable[i][j])
					l--
					status[turnable[i][j]] = false
				} else {
					l--
					status[turnable[i][j]] = false
					break

				}

			}
		}
	}
	for i, v := range wordList {
		if v == beginWord {
			status[i] = true
			break
		}
	}
	t[0] = beginWord
	for i := 0; i < len(wordList); i++ {
		var tem = true
		k := 0
		for ; k < len(beginWord); k++ {
			if wordList[i][k] != beginWord[k] {
				if tem {
					tem = false
				} else {
					break
				}
			}

		}
		if k == len(beginWord) {
			if beginWord != wordList[i] {
				t[1] = wordList[i]
				l = 2
				if i == aim {
					res = res[0:0]
					res = append(res, t[0:2])
					return res
				} else {
					status[i] = true
					f(i)
					status[i] = false
				}

			}

		}

	}
	// var re = make([][]string, len(res))
	// var index = 0
	// for _, v := range res {
	// 	if len(v) == minLen {
	// 		re[index] = v
	// 		index++
	// 	}
	// }
	return res
}

// @lc code=end
