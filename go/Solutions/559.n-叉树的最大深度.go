/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */
package Solutions

// @lc code=start
// type Node struct {
// 	Val      int
// 	Children []*Node
// }

func maxDepth(root *Node_) int {
	if root == nil {
		return 0
	}
	t := 0
	for i := 0; i < len(root.Children); i++ {
		t1 := maxDepth(root.Children[i])
		if t1 > t {
			t = t1
		}
	}
	return t + 1
}

// @lc code=end
