package Solutions

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 0
	max := 0
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
		max = 2
	} else {
		dp[1] = 0
	}

	for i := 2; i < len(s); {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = dp[i-2] + 2
				if dp[i] > max {
					max = dp[i]
				}
				i++

			} else {
				if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
					if i-2-dp[i-1] >= 0 {
						dp[i] = dp[i-1] + 2 + dp[i-2-dp[i-1]]
					} else {
						dp[i] = dp[i-1] + 2
					}

					if dp[i] > max {
						max = dp[i]
					}
				}
				i++
			}
		} else {
			i++
		}
	}
	return max

}

//栈
//func longestValidParentheses(s string) int {
//	l:=len(s)
//	if l==0||l==1{
//		return 0
//	}
//	max:=0
//	stack:=make([]byte,0)
//	a:=make([]int,0)
//	b:=make([]int8,l)
//	sIndex:=0
//	sMax:=0
//	for i:=0;i<l;i++{
//		if s[i]=='('{
//			if sIndex>=sMax{
//				stack=append(stack,s[i])
//				a=append(a,i)
//				sMax++
//			}else{
//				stack[sIndex]=s[i]
//				a[sIndex]=i
//			}
//			sIndex++
//		}else{
//			if sIndex==0{
//			}else{
//				if stack[sIndex-1]=='('{
//					b[a[sIndex-1]],b[i]=1,1
//					sIndex--
//				}
//			}
//		}
//	}
//	tem:=0
//	b= append(b, 0)
//	for i:=0;i<l+1;{
//		if b[i]==1{
//			i+=2
//			tem+=2
//		}else{
//			if tem>max{
//				max=tem
//			}
//			tem=0
//			i++
//		}
//	}
//	return max
//}

// @lc code=end
