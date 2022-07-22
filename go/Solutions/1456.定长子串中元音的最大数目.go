/*
 * @lc app=leetcode.cn id=1456 lang=golang
 *
 * [1456] 定长子串中元音的最大数目
 */
package Solutions

// @lc code=start
func maxVowels(s string, k int) int {
	preVowelCounet := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			preVowelCounet[i+1] = preVowelCounet[i] + 1
		} else {
			preVowelCounet[i+1] = preVowelCounet[i]
		}

	}
	var maxVowels = -1
	for i := 0; i < len(preVowelCounet)-k; i++ {
		t := preVowelCounet[i+k] - preVowelCounet[i]

		if t > maxVowels {
			maxVowels = t
		}
	}
	return maxVowels
}

// @lc code=end
