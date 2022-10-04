package Solutions

func minAddToMakeValid(s string) int {
	var stack = make([]byte, len(s))
	var stackTop = -1
	var ans = 0
	for i := 0; i < len(s); i++ {
		if stackTop == -1 {
			if s[i] == ')' {
				ans++
				continue
			} else {
				stackTop++
				stack[0] = '('
			}
		} else {
			if s[i] == '(' {
				stackTop++
				stack[stackTop] = '('
			} else {
				stackTop--
			}
		}
	}
	ans += stackTop + 1
	return ans
}
