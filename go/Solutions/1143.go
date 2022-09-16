package Solutions

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func longestCommonSubsequence(text1 string, text2 string) int {
	var dp = make([][]int, len(text1))
	for i := range dp {
		dp[i] = make([]int, len(text2))
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
