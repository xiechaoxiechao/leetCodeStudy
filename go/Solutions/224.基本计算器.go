/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

// @lc code=start
package Solutions

import "C"
import "strings"

func calculate(s string) int {
	s = strings.TrimSpace(s)
	flagStack := make([]bool, 100000)
	sTop := 0
	flagStack[0] = true
	var ans = 0
	i := 0
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		ans = ans*10 + int(s[i]-'0')
	}
	for ; i < len(s); i++ {
		if s[i] == '(' {
			if s[i-1] == '(' || s[i-1] == '+' {
				flagStack[sTop+1] = flagStack[sTop]
			}
			flagStack[sTop+1] = !flagStack[sTop]
			sTop++
		} else if s[i] == ')' {
			sTop--

		} else if s[i] >= 0 && s[i] <= '9' {
			var num = 0
			k := i
			for ; k < len(s) && s[k] >= '0' && s[k] <= '9'; k++ {
				num = num*10 + int(s[k]-'0')
			}
			if s[i-1] == '-' {
				if flagStack[sTop] {
					ans -= num
				} else {
					ans += num
				}
			} else {
				if flagStack[sTop] {
					ans += num
				} else {
					ans -= num
				}
			}
			i = k
		}

	}
	return ans
}

// @lc code=end
