/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	height := getTreeHeight(root) - 1
	ans := 0
	for i := 0; i < height; i++ {
		ans += 1 << i
	}
	if height < 0 {
		return 0
	}
	l, r := 0, 1<<height-1
	if exist(root, height, r) {
		return ans + 1<<height
	}
	for l < r {
		mid := l + (r-l)/2
		if exist(root, height, mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return ans + r

}

func getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getTreeHeight(root.Left) + 1
}
func exist(root *TreeNode, height int, index int) bool {
	if height == 0 {
		if root == nil {
			return false
		} else {
			return true
		}
	}
	mid := (1 << (height - 1))
	if index < mid {
		return exist(root.Left, height-1, index)
	} else {
		return exist(root.Right, height-1, index-mid)
	}

}

// @lc code=end
