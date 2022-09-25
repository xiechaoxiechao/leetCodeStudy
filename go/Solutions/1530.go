package Solutions

func countPairs_(root *TreeNode, distance int) int {
	var ans = 0
	dfs_countPa(root, &ans, distance)
	return ans
}

func dfs_countPa(root *TreeNode, ans *int, distance int) []int {
	if root == nil {
		return []int{}
	}
	leftMap := dfs_countPa(root.Left, ans, distance)
	rightMap := dfs_countPa(root.Right, ans, distance)
	if len(leftMap) == 0 && len(rightMap) == 0 {
		return []int{0}
	}
	distance -= 2
	for i := 0; i < len(leftMap); i++ {
		for j := 0; j < len(rightMap); j++ {
			if leftMap[i]+rightMap[j] <= distance {
				*ans++
			} else {
				break
			}
		}
	}
	var sl = make([]int, 0, len(leftMap)+len(rightMap))
	i, j := 0, 0
	for i < len(leftMap) && j < len(rightMap) {
		if leftMap[i] <= rightMap[j] {
			sl = append(sl, leftMap[i]+1)
			i++
		} else {
			sl = append(sl, rightMap[j]+1)
			j++
		}
	}
	for i < len(leftMap) {
		sl = append(sl, leftMap[i]+1)
		i++
	}
	for j < len(rightMap) {
		sl = append(sl, rightMap[j]+1)
		j++
	}
	return sl

}
