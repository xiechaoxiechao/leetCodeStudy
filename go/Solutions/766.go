package Solutions
func isToeplitzMatrix(matrix [][]int) bool {

	for i := 0; i < len(matrix[0]); i++ {
		i1 := i
		k := matrix[0][i]
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i1] != k {
				return false
			}
			i1++
            if i1>=len(matrix[0]){
                break
            }
		}
	}
	for i := 1; i < len(matrix); i++ {
		i1 := i
		k := matrix[i][0]
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i1][j] != k {
				return false
			}
			i1++
            if i1>=len(matrix){
                break
            }
		}

	}
	return true
}