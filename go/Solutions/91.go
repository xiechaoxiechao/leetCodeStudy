package Solutions 

func numDecodings(s string) int {
	var dp = make([]int, len(s)+1)
	if s[0] == '0' {
		return 0
	}
	dp[1] = 1
	dp[0] = 1
	for i := 2; i < len(dp); i++ {
		temNum := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if s[i-1] == '0' && (s[i-2] > '2' || s[i-2] == '0') {
			return 0
		}
		if (temNum > 10 && temNum < 20) || (temNum > 20 && temNum <= 26) {
			dp[i] = dp[i-1] + dp[i-2]
		} else if temNum == 10 || temNum == 20 {
			dp[i] = dp[i-2]
		} else if temNum > 26 {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}
