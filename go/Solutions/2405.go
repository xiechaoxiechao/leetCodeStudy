package Solutions

func partitionString(s string) int {
	mp := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if mp&(1<<(s[i]-'a')) != 0 {
			ans++
			mp = (1 << (s[i] - 'a'))
		} else {
			mp |= (1 << (s[i] - 'a'))
		}
	}
	if mp != 0 {
		ans++
	}
	return ans
}
