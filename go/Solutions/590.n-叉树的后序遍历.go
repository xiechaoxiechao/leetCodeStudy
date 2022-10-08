/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N 叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
package Solutions

func postorder(root *Node_) []int {
	var f func(node *Node_)
	if root == nil {
		return []int{}
	}
	var ans = make([]int, 0, 10000)
	f = func(node *Node_) {
		for i := 0; i < len(node.Children); i++ {
			f(node.Children[i])
		}
		ans = append(ans, node.Val)
	}
	f(root)
	return ans

}

// @lc code=end
