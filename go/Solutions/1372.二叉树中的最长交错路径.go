/*
 * @lc app=leetcode.cn id=1372 lang=golang
 *
 * [1372] 二叉树中的最长交错路径
 */
package Solutions

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	var max = -1
	dfs_long(&max, root)
	return max - 1
}

func dfs_long(max *int, root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l1, _ := dfs_long(max, root.Left)
	_, r2 := dfs_long(max, root.Right)
	l := max_long(r2+1, 1)
	*max = max_long(*max, l)
	r := max_long(l1+1, 1)
	*max = max_long(*max, r)
	return l, r
}
func max_long(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end
