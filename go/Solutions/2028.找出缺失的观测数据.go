/*
 * @lc app=leetcode.cn id=2028 lang=golang
 *
 * [2028] 找出缺失的观测数据
 */

// @lc code=start
package Solutions

func missingRolls(rolls []int, mean int, n int) []int {

	al := (len(rolls) + n) * mean
	for i := 0; i < len(rolls); i++ {
		al -= rolls[i]
	}

	if al > 6*n || al < n {
		return []int{}
	}
	ans := make([]int, n)
	k := al / n
	m := al - k*n
	for i := 0; i < n; i++ {
		if i < m {
			ans[i] = k + 1
		} else {
			ans[i] = k
		}

	}
	return ans

}

// @lc code=end
