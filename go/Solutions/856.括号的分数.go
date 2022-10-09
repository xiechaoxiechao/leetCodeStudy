/*
 * @lc app=leetcode.cn id=856 lang=golang
 *
 * [856] 括号的分数
 */
package Solutions

// @lc code=start
func scoreOfParentheses(s string) int {
	stack := make([]int, len(s)*2)
	sT := -1
	stackI := make([]int, len(s)*2)
	sTI := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			sT++
			stack[sT] = -1
			sTI++
			stackI[sTI] = sT
		} else {
			if stackI[sTI] == sT {
				stack[stackI[sTI]] = 1
				if stackI[sTI]-1 >= 0 && stack[stackI[sTI]-1] != -1 {
					stack[stackI[sTI]-1] += 1
					sT--
				}
				sTI--
			} else {
				stack[stackI[sTI]] = 2 * stack[sT]
				sT--
				if stackI[sTI]-1 >= 0 && stack[stackI[sTI]-1] != -1 {
					stack[stackI[sTI]-1] += stack[stackI[sTI]]
					sT--
				}
				sTI--
			}
		}
	}
	return stack[0]

}

// @lc code=end
