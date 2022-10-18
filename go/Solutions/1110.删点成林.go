/*
 * @lc app=leetcode.cn id=1110 lang=golang
 *
 * [1110] 删点成林
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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var mp = make(map[int]struct{}, len(to_delete))
	for i := 0; i < len(to_delete); i++ {
		mp[to_delete[i]] = struct{}{}
	}
	var dfs func(node *TreeNode, parent *TreeNode, flag bool) []*TreeNode

	dfs = func(node *TreeNode, parent *TreeNode, flag bool) []*TreeNode {
		if node == nil {
			return nil
		}
		if _, ok := mp[node.Val]; ok {
			if parent != nil {
				if flag {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			}
			t1 := dfs(node.Left, nil, true)
			t2 := dfs(node.Right, nil, true)
			return append(t1, t2...)
		} else {
			t1 := dfs(node.Left, node, true)
			t2 := dfs(node.Right, node, false)
			if parent == nil {
				t1 = append(t1, node)
			}
			return append(t1, t2...)
		}
	}
	return dfs(root, nil, true)

}

// @lc code=end
