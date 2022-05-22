/*
 * @lc app=leetcode.cn id=1373 lang=golang
 *
 * [1373] 二叉搜索子树的最大键值和
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

func maxSumBST(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, bool, int, int)
	var max = 0
	dfs = func(root *TreeNode) (int, bool, int, int) {
		if root.Left == nil && root.Right == nil {
			if root.Val > max {
				max = root.Val
			}
			return root.Val, true, root.Val, root.Val
		}
		var sum = root.Val
		var flag = true
		var mi1, ma1 int

		if root.Left != nil {
			t, ok, mi, ma := dfs(root.Left)
			mi1 = mi
			if ok && ma < root.Val {
				sum += t
			} else {
				flag = false
			}
		} else {
			mi1 = root.Val
		}
		if root.Right != nil {
			t, ok, mi, ma := dfs(root.Right)
			ma1 = ma
			if ok && mi > root.Val {
				sum += t
			} else {
				flag = false
			}
		} else {
			ma1 = root.Val
		}
		if flag && max < sum {
			max = sum
		}
		return sum, flag, mi1, ma1
	}
	dfs(root)

	return max
}

// @lc code=end
