/*
 * @lc app=leetcode.cn id=2024 lang=golang
 *
 * [2024] 考试的最大困扰度
 */

// @lc code=start
package Solutions

func maxConsecutiveAnswers(answerKey string, k int) int {
	var max = 1

	var aim byte = 'T'
	for j := 0; j < 2; j++ {
		var kt = k
		if answerKey[0] != aim {
			kt--
		}
		for l, r := 0, 1; r < len(answerKey); r++ {

			if answerKey[r] != aim {
				if kt != 0 {
					kt--
				} else {
					for answerKey[l] == aim {
						l++
					}
					l++
				}
			}
			if max < r-l+1 {
				max = r - l + 1
			}
		}
		aim = 'F'

	}

	return max
}

// @lc code=end
