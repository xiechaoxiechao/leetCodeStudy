package Solutions

func maxLengthBetweenEqualCharacters(s string) int {
	if len(s) == 1  {
		return 0
	}
	var l, r = 0, 0
	for ; l < len(s); r, l = r+1, l+1 {
		for k := len(s) - 1; k > r; k-- {
			if s[l] == s[k] {
				r = k
				break
			}
		}
	}
	return r - l - 1
}
