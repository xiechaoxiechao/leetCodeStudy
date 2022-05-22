package Solutions

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	tem := make([]int, k)
	var cp []int
	var recursive func(int)
	recursive = func(index int) {
		i := tem[index-1] + 1
		if k-index < n-i+1 {
			return
		}
		for ; i <= n; i++ {
			tem[index] = i
			if index == k-1 {
				cp = make([]int, k)
				copy(cp, tem)
				res = append(res, cp)
			} else {
				recursive(index + 1)
			}
		}
	}

	for i := 1; i <= n; i++ {
		tem[0] = i
		if k == 1 {
			cp = make([]int, k)
			copy(cp, tem)
			res = append(res, cp)
		} else {
			recursive(1)
		}
	}
	return res
}

// @lc code=end
