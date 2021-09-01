package Solutions

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	arr := make([]int, 0)
	tem := 0
	index := 0
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
	//fmt.Println(candidates)
	res = ana1(candidates, target, arr, res, 0)
	return res
}

func ana1(c []int, target int, arr []int, res [][]int, begin int) [][]int {
	for i := begin; i < len(c); i++ {
		if c[i] == target {
			arr = append(arr, c[i])
			res = append(res, make([]int, len(arr)))
			copy(res[len(res)-1], arr)
			arr = arr[:len(arr)-1]
		} else {
			target -= c[i]
			if target < 0 {
				return res
			}
			arr = append(arr, c[i])
			res = ana1(c, target, arr, res, i)
			arr = arr[:len(arr)-1]
			target += c[i]
		}
	}
	return res
}

// @lc code=end
