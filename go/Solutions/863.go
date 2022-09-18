package Solutions


func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var parentMap = map[int]*TreeNode{}
	var dfs1 func(*TreeNode)
	dfs1 = func(node *TreeNode) {
		if node.Left != nil {
			parentMap[node.Left.Val] = node
			dfs1(node.Left)
		}
		if node.Right != nil {
			parentMap[node.Right.Val] = node
			dfs1(node.Right)
		}
	}
	var dfs2 func(*TreeNode, int, *TreeNode)
	var ans = make([]int, 0)
	dfs2 = func(node *TreeNode, distance int, from *TreeNode) {
		if distance+1 == k {
			ans = append(ans, node.Val)
		}
		if p, ok := parentMap[node.Val]; ok && p != from {
			dfs2(p, distance+1, node)
		}
		if node.Left != nil && node.Left != from {
			dfs2(node.Left, distance+1, node)
		}
		if node.Right != nil && node.Right != from {
			dfs2(node.Right, distance, node)
		}
	}
	dfs2(target, 0, nil)
	return ans
}
