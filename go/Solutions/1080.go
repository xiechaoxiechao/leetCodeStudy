package Solutions

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	if dfs_suff(root, limit, 0) {
		return root
	} else {
		return nil
	}
}
func dfs_suff(node *TreeNode, limit int, sum int) bool {
	var flag1, flag2 bool
	if node.Left != nil {
		if !dfs_suff(node.Left, limit, sum+node.Val) {
			node.Left = nil
			flag1 = true
		}
	}
	if node.Right != nil {
		if !dfs_suff(node.Right, limit, sum+node.Val) {
			flag2 = true
			node.Right = nil
		}
	}
	if node.Left == nil && node.Right == nil {
		if sum+node.Val < limit || (flag1 || flag2) {
			return false
		}
	}
	return true
}
