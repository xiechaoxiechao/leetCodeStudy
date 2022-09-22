package Solutions


func canFormArray(arr []int, pieces [][]int) bool {
	var mp = make(map[int]int)
	for i, v := range pieces {
		mp[v[0]] = i
	}
	for i := 0; i < len(arr); {
		a, ok := mp[arr[i]]
		if !ok {
			return false
		}
		for j := 0; j < len(pieces[a]); j++ {
			if arr[i] != pieces[a][j] {
				return false
			}
			i++
		}
		i += len(pieces[a])
	}
	return true
}