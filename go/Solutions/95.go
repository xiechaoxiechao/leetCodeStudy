package Solutions
func generateTrees(n int) []*TreeNode {
	var trees = make([][][]*TreeNode, n)
	for i := range trees {
		trees[i] = make([][]*TreeNode, n)
		for j := range trees[i] {
			trees[i][j] = make([]*TreeNode, 0, 8)
		}
	}
	for i := 0; i < n; i++ {
		trees[i][i] = append(trees[i][i], &TreeNode{
			i + 1,
			nil,
			nil,
		})
	}
	for i := 1; i < n; i++ {
		for j, i1 := 0, i; i1 < n; j, i1 = j+1, i1+1 {
			for k := j; k <= i1; k++ {
				if k-1 < j {
					for _, v1 := range trees[k+1][i1] {
						trees[j][i1] = append(trees[j][i1], &TreeNode{
							k + 1,
							nil,
							v1,
						})
					}
					continue
				}
				if k+1 > i1 {
					for _, v := range trees[j][k-1] {
						trees[j][i1] = append(trees[j][i1], &TreeNode{
							k + 1,
							v,
							nil,
						})

					}
					continue
				}
				for _, v := range trees[j][k-1] {
					for _, v1 := range trees[k+1][i1] {
						trees[j][i1] = append(trees[j][i1], &TreeNode{
							k + 1,
							v,
							v1,
						})
					}

				}
			}
		}
	}
	return trees[0][n-1]
}
