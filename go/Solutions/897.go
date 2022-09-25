package Solutions

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head *TreeNode = nil
	_, head = dfs_increa(root)
	return head
}

func dfs_increa(node *TreeNode) (*TreeNode, *TreeNode) {
	var end *TreeNode
	var head *TreeNode
	if node.Left != nil {
		end, head = dfs_increa(node.Left)
		end.Right = node
		end = node
	} else {
		head = node
		end = node
	}

	if node.Right != nil {
		t, h := dfs_increa(node.Right)
		end.Right = h
		end = t
	}
	return end, head

}
