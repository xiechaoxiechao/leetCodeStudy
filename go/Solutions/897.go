package Solutions 



func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head *TreeNode = nil
	_, head = dfs(root)
	return head
}

func dfs(node *TreeNode) (*TreeNode, *TreeNode) {
	var end *TreeNode
	var head *TreeNode
	if node.Left != nil {
		end, head = dfs(node.Left)
		end.Right = node
		end = node
	} else {
		head = node
		end = node
	}

	if node.Right != nil {
		t, h := dfs(node.Right)
		end.Right = h
		end = t
	}
	return end, head

}