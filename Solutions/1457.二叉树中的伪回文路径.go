/*
 * @lc app=leetcode.cn id=1457 lang=golang
 *
 * [1457] 二叉树中的伪回文路径
 */

// @lc code=start
package Solutions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	var ans = 0
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, tem int) {
		var n = tem ^ (1 << node.Val)
		if node.Left == nil && node.Right == nil {
			if n == 0 || n&(n-1) == 0 {
				ans++
			}
		}
		if node.Left != nil {
			dfs(node.Left, n)
		}
		if node.Right != nil {
			dfs(node.Right, n)
		}

	}
	dfs(root, 0)
	return ans

}

// @lc code=end
