/*
 * @lc app=leetcode.cn id=2100 lang=golang
 *
 * [2100] 适合打劫银行的日子
 */

// @lc code=start
package Solutions

func goodDaysToRobBank(security []int, time int) []int {
	if time == 0 {
		for i := 0; i < len(security); i++ {
			security[i] = i
		}
		return security
	}
	var back = make([]int, len(security))
	var res = make([]int, 0)
	var last = 0
	for i := len(security) - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			last++
			if last >= time {
				back[i] = 1
			}
		} else {
			last = 0
		}
	}
	last = 0
	for i := 1; i < len(security); i++ {
		if security[i] <= security[i-1] {
			last++
			if back[i] == 1 && last >= time {
				res = append(res, i)
			}
		} else {
			last = 0
		}
	}

	return res
}

// @lc code=end
