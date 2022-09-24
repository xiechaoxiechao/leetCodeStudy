package Solutions

func shiftingLetters(s string, shifts [][]int) string {
	var diff = make([]int, len(s)+1)
	for i := 0; i < len(shifts); i++ {
		if shifts[i][2] == 1 {
			diff[shifts[i][0]] += 1
			diff[shifts[i][1]+1] -= 1
		} else {
			diff[shifts[i][0]] -= 1
			diff[shifts[i][1]+1] += 1
		}
       
	}
  
	ans := make([]byte,len(s))
    t:=(26+diff[0]%26)%26
	ans[0] = (s[0]-'a' +byte(t))%26+'a'
	for i := 1; i < len(s); i++ {
        t+=diff[i]
		ans[i] = (s[i]-'a' +26+byte(t%26))%26+'a'
	}
	return string(ans)
}
