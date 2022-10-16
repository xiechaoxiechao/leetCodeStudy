package Solutions


func smallestSubsequence(s string) string {
	var mp = [26]int{}
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
	}
	var mps = [26]int{}
	stack := make([]byte, len(s))
	st := 0
	stack[0] = s[0]
	mps[s[0]-'a'] = 1
	for i := 1; i < len(s); i++ {
		if mps[s[i]-'a'] == 1 {
             mp[s[i]-'a']--
			continue
		}
		for stack[st] > s[i] {
			if mp[stack[st]-'a'] > 1 {
                mps[stack[st]-'a']=0
                mp[stack[st]-'a']--
				st--
                if st<0{
                    break
                }
			} else {
				break
			}
		}
		st++
		stack[st] = s[i]
		mps[s[i]-'a'] = 1
	}
	return string(stack[:st+1])

}