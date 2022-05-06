/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package Solutions

import "C"

func findTarget(root *TreeNode, k int) bool {
	var mp = make(map[int]struct{})
	var bfs func(nod *TreeNode) bool

	bfs = func(nod *TreeNode) bool {
		tem := k - nod.Val
		if _, ok := mp[tem]; ok {
			return true
		}
		mp[nod.Val] = struct{}{}
		if nod.Left != nil && bfs(nod.Left) == true {
			return true
		}
		if nod.Right != nil && bfs(nod.Right) == true {
			return true
		}
		return false
	}
	return bfs(root)
}

// @lc code=end
