package Solutions


func findClosest(words []string, word1 string, word2 string) int {
	var mp = make(map[string][]int)
	for i, v := range words {
		if v1, ok := mp[v]; ok {
			mp[v] = append(v1, i)
		} else {
			mp[v] = []int{i}
		}
	}
	inds1 := mp[word1]
	inds2 := mp[word2]
	var minDis = 1<<32 - 1
	for i, j := 0, 0; i < len(inds1) && j < len(inds2); {
		t := inds1[i] - inds2[j]
		if t < 0 {
			i++
			t = -t
			if t < minDis {
				minDis = t
			}
			continue
		} else {
			j++
			if t < minDis {
				minDis = t
			}
			continue
		}
	}
	return minDis

}
