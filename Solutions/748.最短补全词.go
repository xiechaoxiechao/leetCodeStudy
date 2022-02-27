/*
 * @lc app=leetcode.cn id=748 lang=golang
 *
 * [748] 最短补全词
 */
package Solutions

// @lc code=start
func shortestCompletingWord(licensePlate string, words []string) string {
	var Acount = [26]uint8{}
	var count = [26]uint8{}
	var l = 0
	for _, v := range licensePlate {
		if v > 96 {
			Acount[v-'a']++
			l++
		} else if v > 64 {
			Acount[v-'A']++
			l++
		}
	}
	var minLen = 1001
	var res = -1
	for i, v := range words {
		if len(v) < l {
			continue
		}
		for _, c := range v {
			if c > 96 {
				count[c-'a']++
			} else if c > 64 {
				count[c-'A']++
			}

		}

		flag := true
		for k := 0; k < 26; k++ {
			if count[k] < Acount[k] {
				flag = false
			}
			count[k] = 0
		}
		if flag {
			if minLen > len(v) {
				res = i
				minLen = len(v)
			}
		}

	}
	return words[res]
}

// @lc code=end
