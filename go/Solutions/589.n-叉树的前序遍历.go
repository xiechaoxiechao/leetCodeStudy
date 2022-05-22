/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start

package Solutions

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res = make([]int, 1000)
	var f func(node *Node)
	f = func(node *Node) {
		res = append(res, node.Val)
		for i := 0; i < len(node.Children); i++ {
			f(node.Children[i])
		}
	}
	return res
}

// @lc code=end
