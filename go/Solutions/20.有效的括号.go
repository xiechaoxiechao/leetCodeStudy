package Solutions

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	length1 := len(s)
	if length1%2 != 0 {
		return false
	}
	p := make([]byte, 0, 1)
	nowLen := -1
	for i := 0; i < length1; i++ {

		if nowLen < 0 {
			p = append(p, s[i])
			nowLen++
			continue
		}

		if s[i] != p[nowLen]+1 && s[i] != p[nowLen]+2 {
			p = append(p, s[i])
			nowLen++
		} else {
			p = p[0:nowLen]
			nowLen--
		}
	}
	if len(p) == 0 {
		return true
	}

	return false
}

// @lc code=end
