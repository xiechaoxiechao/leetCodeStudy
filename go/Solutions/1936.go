package Solutions

func addRungs(rungs []int, dist int) int {
	var ans = 0
	var now = 0
	for i := 0; i < len(rungs); {
		distance := rungs[i] - now
		if distance > dist {
			t := (rungs[i] - now) / dist
			ans += t
			now = now + dist*t
			if now == rungs[i] {
				ans--
			}
		} else {
			now = rungs[i]
			i++
		}
	}
	return ans
}