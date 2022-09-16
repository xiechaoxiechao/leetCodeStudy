package Solutions 
func isInterleave(s1 string, s2 string, s3 string) bool {
	var dp = make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
    if s1==""{
        if s2!=s3{
            return false
        }
        return true
    }
    if s2==""{
        if s1!=s3{
            return false
        }
        return true
    }
    if len(s1)+len(s2)!=len(s3){
        return false
    }
	for i := range s1 {
		if s1[i] == s3[i] {
			dp[i+1][0] = true
		} else {
			break
		}
	}
	for i := range s2 {
		if s2[i] == s3[i] {
			dp[0][i+1] = true
		} else {
			break
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s3[i+j-1] == s1[i-1] {
				if dp[i-1][j] {
					dp[i][j] = true
					continue
				}
			}
			if s3[i+j-1] == s2[j-1] {
				if dp[i][j-1] {
					dp[i][j] = true
					continue
				}
			}

		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
