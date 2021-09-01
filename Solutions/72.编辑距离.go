package Solutions

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	var dp = make([][]int, len(word1))
	dp[0] = make([]int, len(word2))
	if word1[0] == word2[0] {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	minAB := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	var i int
	for i = 1; i < len(word1); i++ {
		dp[i] = append(dp[i], make([]int, len(word2))...)
		if word1[i] == word2[0] {
			dp[i][0] = i
		} else {
			dp[i][0] = dp[i-1][0] + 1
		}

	}
	for i = 1; i < len(word2); i++ {
		if word1[0] == word2[i] {
			dp[0][i] = i
		} else {
			dp[0][i] = dp[0][i-1] + 1
		}
	}
	for i = 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minAB(dp[i][j-1], minAB(dp[i-1][j-1], dp[i-1][j])) + 1
			}
		}
	}
	return dp[len(word1)-1][len(word2)-1]
}

// @lc code=end
