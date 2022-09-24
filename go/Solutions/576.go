package Solutions

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var dp = make([][][]int, m+2)
	for i := 0; i < m+2; i++ {
		dp[i] = make([][]int, n+2)
		for j := 0; j < n+2; j++ {
			dp[i][j] = make([]int, maxMove+1)
			for k := 0; k < maxMove+1; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	for i := 0; i < n+2; i++ {
		for j := 0; j < maxMove+1; j++ {
			dp[0][i][j] = 1
			dp[m+1][i][j] = 1
		}

	}
	for i := 0; i < m+2; i++ {
		for j := 0; j < maxMove+1; j++ {
			dp[i][0][j] = 1
			dp[i][n+1][j] = 1
		}
	}

	var dfs func(i, j, step int) int
	dfs = func(i, j, step int) int {
		if step == -1 {
			return 0
		}
		if dp[i][j][step] != -1 {
			return dp[i][j][step]
		}
		if step == 0 {
			dp[i][j][step] = 0
			return 0
		}
		t := dfs(i-1, j, step-1) + dfs(i, j-1, step-1) + dfs(i+1, j, step-1) + dfs(i, j+1, step-1)
		dp[i][j][step] = t% 1000000007
		return dp[i][j][step]

	}

	return dfs(startRow+1, startColumn+1, maxMove) % 1000000007

}