/*
 * @lc app=leetcode.cn id=1106 lang=golang
 *
 * [1106] 解析布尔表达式
 */

// @lc code=start
package Solutions

func countEval(s string, result int) int {
	var l = len(s)
	var num = make([]bool, 0, l)
	var ope = make([]int, 0, l)
	for i := 0; i < l; i += 2 {
		if s[i] == '1' {
			num = append(num, true)
		} else {
			num = append(num, false)
		}
	}
	for i := 1; i < l; i += 2 {
		switch s[i] {
		case '&':
			ope = append(ope, 1)
		case '|':
			ope = append(ope, 2)
		case '^':
			ope = append(ope, 3)
		}
	}

	var dpOne = make([][]int, len(num))
	var dpZero = make([][]int, len(num))
	for i := 0; i < len(num); i++ {
		dpOne[i] = make([]int, len(num))
		dpZero[i] = make([]int, len(num))
		for j := 0; j < len(num); j++ {
			dpOne[i][j] = -1
			dpZero[i][j] = -1
		}
		if num[i] {
			dpOne[i][i] = 1
			dpZero[i][i] = 0
		} else {
			dpOne[i][i] = 0
			dpZero[i][i] = 1
		}
	}
	var _c func(int, int, bool) int = nil
	_c = func(l, r int, aim bool) int {
		if aim {
			if dpOne[l][r] != -1 {
				return dpOne[l][r]
			}
		} else {
			if dpZero[l][r] != -1 {
				return dpZero[l][r]
			}
		}

		var time = 0
		for i := l; i < r; i++ {
			if ope[i] == 1 {
				if aim {
					time += _c(l, i, aim) * _c(i+1, r, aim)
				} else {
					time += _c(l, i, aim) * _c(i+1, r, aim)
					time += _c(l, i, !aim) * _c(i+1, r, aim)
					time += _c(l, i, aim) * _c(i+1, r, !aim)
				}

			} else if ope[i] == 2 {
				if aim {
					time += _c(l, i, !aim) * _c(i+1, r, aim)
					time += _c(l, i, aim) * _c(i+1, r, !aim)
					time += _c(l, i, aim) * _c(i+1, r, aim)
				} else {
					time += _c(l, i, aim) * _c(i+1, r, aim)
				}

			} else {
				if aim {
					time += _c(l, i, aim) * _c(i+1, r, !aim)
					time += _c(l, i, !aim) * _c(i+1, r, aim)
				} else {
					time += _c(l, i, aim) * _c(i+1, r, aim)
					time += _c(l, i, !aim) * _c(i+1, r, !aim)
				}

			}
		}
		if aim {
			dpOne[l][r] = time
		} else {
			dpZero[l][r] = time
		}

		return time
	}
	if result == 1 {
		return _c(0, len(num)-1, true)
	} else {
		return _c(0, len(num)-1, false)
	}

}
func ac(a int) int {
	return (a+1)*a/2 - a + 1
}

// @lc code=end
