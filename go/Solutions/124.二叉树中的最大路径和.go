/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
package Solutions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var ans = root.Val
	var getMax func(*TreeNode) int
	getMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		var l = 0
		var r = 0
		if node.Left != nil {
			l = getMax(node.Left)
		}
		if node.Right != nil {
			r = getMax(node.Right)
		}
		var max = 0
		if l > r {
			max = l
		} else {
			max = r
		}
		tem := node.Val
		if l > 0 {
			tem += l
		}
		if r > 0 {
			tem += r
		}

		if tem > ans {
			ans = tem
		}
		if max > 0 {
			return max + node.Val
		}
		return node.Val

	}

	getMax(root)
	return ans
}

// @lc code=end
