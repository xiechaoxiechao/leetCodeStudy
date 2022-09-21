package Solutions

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var ans = 0
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}
	return ans
}
