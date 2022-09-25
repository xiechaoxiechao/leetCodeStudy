package Solutions

import "sort"

func sortString(s string) string {
	mp := map[byte]int{}
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}
	var byteKeys = make([]byte, 0, len(mp))
	for k := range mp {
		byteKeys = append(byteKeys, k)
	}
	sort.Slice(byteKeys, func(i, j int) bool {
		return byteKeys[i] < byteKeys[j]
	})
	var ans = make([]byte, 0, len(s))
	for len(mp) != 0 {
		for i := 0; i < len(byteKeys); i++ {
			if v, ok := mp[byteKeys[i]]; ok {
				ans = append(ans, byteKeys[i])
				if v--; v != 0 {
					mp[byteKeys[i]] = v
				} else {
					delete(mp, byteKeys[i])
				}
			}
		}
		for i := len(byteKeys) - 1; i > -1; i-- {
			if v, ok := mp[byteKeys[i]]; ok {
				ans = append(ans, byteKeys[i])
				if v--; v != 0 {
					mp[byteKeys[i]] = v
				} else {
					delete(mp, byteKeys[i])
				}
			}
		}

	}
	return string(ans)

}
