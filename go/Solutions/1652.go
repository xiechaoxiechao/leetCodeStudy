package Solutions

func decrypt(code []int, k int) []int {
	var t = make([]int, 0, len(code)*3)
	t = append(t, code...)
	t = append(t, code...)
	t = append(t, code...)
	var ans = make([]int, len(code))
	for i := 0; i < len(code); i++ {
		if k == 0 {
			ans[i] = 0
		} else if k < 0 {
			n := i + len(code)
			sum := 0
			for j := 0; j > k; j-- {
				n--
				sum += t[n]
			}
			ans[i] = sum
		} else {
			n := i + len(code)
			sum := 0
			for j := 0; j < k; j++ {
				n++
				sum += t[n]
			}
			ans[i] = sum
		}
	}
	return ans
}
