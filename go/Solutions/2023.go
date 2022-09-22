package Solutions

func numOfPairs(nums []string, target string) int {
	var preMp = make(map[byte][]int, len(nums))
	var ans = 0
	for i, v := range nums {
		if len(v) >= len(target) {
			continue
		}
		if s, ok := preMp[v[0]]; ok {
			preMp[v[0]] = append(s, i)
		} else {
			preMp[v[0]] = []int{i}
		}
	}
	first := preMp[target[0]]
	for i := 0; i < len(first); i++ {
		second := preMp[target[len(nums[first[i]])]]
		for j := 0; j < len(second); j++ {
			if first[i] == second[j] {
				continue
			}
			if nums[first[i]]+nums[second[j]] == target {
				ans++
			}
		}

	}
	return ans
}
