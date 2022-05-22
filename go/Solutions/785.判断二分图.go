/*
 * @lc app=leetcode.cn id=785 lang=golang
 *
 * [785] 判断二分图
 */
package Solutions

// @lc code=start
func isBipartite(graph [][]int) bool {
	var status = make([]int, len(graph))
	var dfs func(int, bool) bool

	dfs = func(node int, flag bool) bool {
		if flag {
			for i := 0; i < len(graph[node]); i++ {
				if status[graph[node][i]] == 0 {
					status[graph[node][i]] = 1
					if !dfs(graph[node][i], false) {
						return false
					}
				} else if status[graph[node][i]] == 2 {
					return false
				}
			}

		} else {
			for i := 0; i < len(graph[node]); i++ {
				if status[graph[node][i]] == 0 {
					status[graph[node][i]] = 2
					if !dfs(graph[node][i], true) {
						return false
					}
				} else if status[graph[node][i]] == 1 {
					return false
				}
			}
		}

		return true

	}
	for i := 0; i < len(graph); i++ {
		if status[i] == 0 {
			status[i] = 1
			if !dfs(i, false) {
				return false
			}
		}

	}
	return true

}

// @lc code=end
