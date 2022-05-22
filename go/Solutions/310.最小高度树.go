/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */

// @lc code=start
package Solutions

func findMinHeightTrees(n int, edges [][]int) []int {

	var maxDis = 0
	var y = -1
	var dfs func(root int, dep int, in int)
	var dfs1 func(root int, dep int, in int)

	var mps = make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		mps[i] = make(map[int]struct{})
	}
	for i := 0; i < len(edges); i++ {
		mps[edges[i][0]][edges[i][1]] = struct{}{}
		mps[edges[i][1]][edges[i][0]] = struct{}{}

	}
	var stack = make([]int, n)
	var stackTop = 0
	var maxStackTop = 0
	var path = make([]int, n)
	dfs = func(root, dep, in int) {
		stack[stackTop] = root
		stackTop++
		dep++
		for key := range mps[root] {
			if key == in {
				continue
			}
			dfs(key, dep, root)
		}

		if dep > maxDis {
			copy(path, stack[:stackTop])
			maxStackTop = stackTop - 1
			maxDis = dep
		}
		stackTop--
	}
	dfs1 = func(root, dep, in int) {
		dep++
		for key := range mps[root] {
			if key == in {
				continue
			}
			dfs1(key, dep, root)
		}

		if dep > maxDis {
			maxDis = dep
			y = root
		}
	}
	dfs1(0, 0, 0)
	maxDis = 0
	dfs(y, 0, y)
	if (maxStackTop)%2 == 1 {
		return path[maxStackTop/2 : maxStackTop/2+2]
	} else {
		return path[maxStackTop/2 : maxStackTop/2+1]
	}

}

// @lc code=end
