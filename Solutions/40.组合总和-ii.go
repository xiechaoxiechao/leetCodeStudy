package Solutions

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	arr := make([]int, 0)
	tem := 0
	index := 0
	ha := make(map[int]int)
	for i := 0; i < len(candidates); i++ {
		tem = candidates[i]
		index = i
		for j := i; j < len(candidates); j++ {
			if tem > candidates[j] {
				tem = candidates[j]
				index = j
			}
		}
		candidates[i], candidates[index] = candidates[index], candidates[i]
	}
	tem = candidates[0]
	ha[tem] = 1
	for i := 1; i < len(candidates); i++ {
		if tem == candidates[i] {
			ha[tem]++
		} else {
			tem = candidates[i]
			ha[tem]++
		}
	}
	res = ana(candidates, target, arr, res, 0, ha)
	return res
}

func ana(c []int, target int, arr []int, res [][]int, begin int, ha map[int]int) [][]int {
	for i := begin; i < len(c); i += ha[c[i]] {
		for j := 0; j < ha[c[i]]; j++ {
			target -= (j + 1) * c[i]
			if target < 0 {
				target += (j + 1) * c[i]
				break
			}
			for k := 0; k <= j; k++ {
				arr = append(arr, c[i])
			}
			if target == 0 {
				res = append(res, make([]int, len(arr)))
				copy(res[len(res)-1], arr)
				arr = arr[:len(arr)-1-j]
				target += (j + 1) * c[i]
				break
			} else {
				res = ana(c, target, arr, res, i+ha[c[i]], ha)
				arr = arr[:len(arr)-1-j]
				target += (j + 1) * c[i]
			}

		}
	}
	return res
}

// @lc code=end
