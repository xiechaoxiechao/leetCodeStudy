package Solutions

func numTrees(n int) int {

	var trees = make([][]int, n)
	for i := range trees {
		trees[i] = make([]int, n)
		trees[i][i] = 1

	}
	for i := 1; i < n; i++ {
		for j, i1 := 0, i; i1 < n; j, i1 = j+1, i1+1 {
			for k := j; k <= i1; k++ {
				if k-1 < j {
					trees[j][i1] += trees[k+1][i1]
					continue
				}
				if k+1 > i1 {
					trees[j][i1] +=trees[j][k-1]
					continue
				}

				trees[j][i1] += trees[j][k-1] * trees[k+1][i1]

			}
		}
	}
	return trees[0][n-1]

}
