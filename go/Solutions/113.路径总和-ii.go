/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	return dfs_ps(root, 1, targetSum, targetSum)
}

func dfs_ps(node *TreeNode, depth int, target int, sum int) [][]int {
	if node.Left == nil && node.Right == nil {
		if node.Val != sum {
			return nil
		}
		sl := make([][]int, 0)
		sl = append(sl, make([]int, depth))
		sl[0][depth-1] = node.Val
		return sl
	}
	var sl [][]int
	if node.Left != nil {
		sl = dfs_ps(node.Left, depth+1, target, sum-node.Val)
		if sl != nil {
			for i := 0; i < len(sl); i++ {
				sl[i][depth-1] = node.Val
			}
		}
	}
	if node.Right != nil {
		sl2 := dfs_ps(node.Right, depth+1, target, sum-node.Val)
		if sl2 != nil {
			for i := 0; i < len(sl2); i++ {
				sl2[i][depth-1] = node.Val
			}
			if sl == nil {
				sl = sl2
			} else {
				sl = append(sl, sl2...)
			}

		}
	}
	return sl

}

// @lc code=end
