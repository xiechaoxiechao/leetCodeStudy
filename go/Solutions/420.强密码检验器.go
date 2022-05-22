/*
 * @lc app=leetcode.cn id=420 lang=golang
 *
 * [420] 强密码检验器
 */

// @lc code=start
package Solutions

func strongPasswordChecker(password string) int {
	var count = [3]int{}
	for i := 0; i < len(password); i++ {
		switch {
		case password[i] >= '0' && password[i] <= '9':
			count[0]++
			break
		case password[i] >= 'A' && password[i] <= 'Z':
			count[1]++
			break
		case password[i] >= 'a' && password[i] <= 'z':
			count[2]++
			break
		}
	}
	l := len(password)
	step := 0
	for i := 0; i < len(password); {
		a := password[i]
		c := 1
		for i++; i < len(password) && password[i] == a; i++ {
			c++
		}

		for c >= 3 {
			k := c - 2
			if l > 20 {
				if k > (l - 20) {
					l = 20
					step += (l - 20)
					c -= (l - 20)
				} else {
					step += k
					l -= k
					c = 2
				}
			} else if l < 6 {
				st := (c - 1) / 2
				if st+l < 20 {
					step += st
					c = 2
					for j := 0; j < st; j++ {
						for p := 0; p < 3; p++ {
							if count[p] == 0 {
								count[p]++
							}
						}
					}
				} else {

				}
			}
		}

	}
	if l < 6 {
		step += (6 - l)
	} else if l > 20 {
		step += (l - 20)
	}
	return step
}

// @lc code=end
