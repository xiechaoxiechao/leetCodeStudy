package Solutions

func levelOrder(root *TreeNode) []int {
	nodeArr := make([]*TreeNode, 0, 8)
	ans := make([]int, 0, 8)
	if root == nil {
		return []int{}
	} else {
		nodeArr = append(nodeArr, root)
	}
	for i := 0; i < len(nodeArr); i++ {
		if nodeArr[i].Left != nil {
			nodeArr = append(nodeArr, nodeArr[i].Left)
		}
		if nodeArr[i].Right != nil {
			nodeArr = append(nodeArr, nodeArr[i].Right)
		}
		ans = append(ans, nodeArr[i].Val)
	}
	return ans
}