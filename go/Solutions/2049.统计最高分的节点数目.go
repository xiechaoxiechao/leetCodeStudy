/*
 * @lc app=leetcode.cn id=2049 lang=golang
 *
 * [2049] 统计最高分的节点数目
 */

// @lc code=start
package Solutions

func countHighestScoreNodes(parents []int) int {
	var childNums = make([]int, len(parents))
	var child = make([][]int, len(parents))
	for i := 0; i < len(child); i++ {
		child[i] = make([]int, 0, 2)
	}
	for i := 1; i < len(parents); i++ {
		child[parents[i]] = append(child[parents[i]], i)
	}
	var dfs func(int) int
	dfs = func(i int) int {
		var ans = 1
		for _, v := range child[i] {
			ans += dfs(v)

		}
		childNums[i] = ans
		return ans
	}
	childNums[0] = dfs(0)
	var max int64 = 1
	var nums = 0
	for i := 0; i < len(parents); i++ {
		var tem int64 = 1
		if a := len(parents) - childNums[i]; a != 0 {
			tem *= int64(a)
		}
		for _, v := range child[i] {
			tem *= int64(childNums[v])
		}
		if max < tem {
			max = tem
			nums = 1
		} else if max == tem {
			nums++
		}

	}
	return nums
}

// @lc code=end
