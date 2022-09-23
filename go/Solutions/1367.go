package Solutions

func isSubPath(head *ListNode, root *TreeNode) bool {
	return match(head, root, head) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func match(head *ListNode, root *TreeNode, front *ListNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val == root.Val {
		if match(head.Next, root.Left, front) || match(head.Next, root.Right, front) {
			return true
		}

	}
	return false
}
