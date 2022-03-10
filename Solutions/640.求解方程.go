/*
 * @lc app=leetcode.cn id=640 lang=golang
 *
 * [640] 求解方程
 */

// @lc code=start
package Solutions

import "strconv"

func solveEquation(equation string) string {
	var lNum = 0
	var lx = 0
	var flag = true
	i := 0
	for equation[i] != '=' {
		if equation[i] == '+' {
			flag = true
			i++
		}
		if equation[i] == '-' {
			flag = false
			i++
		}
		if equation[i] == 'x' {
			if flag {
				lx++
			} else {
				lx--
			}
			i++
		} else {
			var l = i
			for equation[i] > 46 {
				i++
			}
			r := i
			if flag {
				tem, _ := strconv.Atoi(equation[l:r])
				lNum += tem
			} else {
				tem, _ := strconv.Atoi(equation[l:r])
				lNum -= tem
			}
		}

	}
	flag = true
	i++
	for i < len(equation) {
		if equation[i] == '-' {
			flag = false
			i++
		}
		if equation[i] == '+' {
			flag = true
			i++
		}
		if equation[i] == 'x' {
			if flag {
				lx -= 1
			} else {
				lx += 1
			}
			i++

		} else {
			var l = i
			for i < len(equation) && equation[i] > 46 {
				i++
			}
			r := i
			if flag {
				tem, _ := strconv.Atoi(equation[l:r])
				lNum -= tem
			} else {
				tem, _ := strconv.Atoi(equation[l:r])
				lNum += tem
			}
		}

	}
	if lx == 0 && lNum == 0 {
		return "Infinite solutions"
	}
	if lx == 0 {
		return "No solution"
	}
	return "x=" + strconv.Itoa(-lNum/lx)

}

// @lc code=end
